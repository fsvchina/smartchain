3 of the voting power has upgraded, the blockchain will immediately
resume the consensus mechanism. If the majority of operators add a custom `do-upgrade` script, this should
be a matter of minutes and not even require them to be awake at that time.

Integrating With An App

Setup an upgrade Keeper for the app and then define a BeginBlocker that calls the upgrade
keeper's BeginBlocker method:
    func (app *myApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
    	app.upgradeKeeper.BeginBlocker(ctx, req)
    	return abci.ResponseBeginBlock{}
    }

The app must then integrate the upgrade keeper with its governance module as appropriate. The governance module
should call ScheduleUpgrade to schedule an upgrade and ClearUpgradePlan to cancel a pending upgrade.

Performing Upgrades

Upgrades can be scheduled at a predefined block height. Once this block height is reached, the
existing software will cease to process ABCI messages and a new version with code that handles the upgrade must be deployed.
All upgrades are coordinated by a unique upgrade name that cannot be reused on the same blockchain. In order for the upgrade
module to know that the upgrade has been safely applied, a handler with the name of the upgrade must be installed.
Here is an example handler for an upgrade named "my-fancy-upgrade":
	app.upgradeKeeper.SetUpgradeHandler("my-fancy-upgrade", func(ctx sdk.Context, plan upgrade.Plan) {

	})

This upgrade handler performs the dual function of alerting the upgrade module that the named upgrade has been applied,
as well as providing the opportunity for the upgraded software to perform any necessary state migrations. Both the halt
(with the old binary) and applying the migration (with the new binary) are enforced in the state machine. Actually
switching the binaries is an ops task and not handled inside the sdk / abci app.

Here is a sample code to set store migrations with an upgrade:


	app.UpgradeKeeper.SetUpgradeHandler("my-fancy-upgrade",  func(ctx sdk.Context, plan upgrade.Plan) {

	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {

	}

	if upgradeInfo.Name == "my-fancy-upgrade" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{
			Renamed: []store.StoreRename{{
				OldKey: "foo",
				NewKey: "bar",
			}},
			Deleted: []string{},
		}


		app.SetStoreLoader(upgrade.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

Halt Behavior

Before halting the ABCI state machine in the BeginBlocker method, the upgrade module will log an error
that looks like:
	UPGRADE "<Name>" NEEDED at height <NNNN>: <Info>
where Name are Info are the values of the respective fields on the upgrade Plan.

To perform the actual halt of the blockchain, the upgrade keeper simply panics which prevents the ABCI state machine
from proceeding but doesn't actually exit the process. Exiting the process can cause issues for other nodes that start
to lose connectivity with the exiting nodes, thus this module prefers to just halt but not exit.

Automation and Plan.Info

We have deprecated calling out to scripts, instead with propose https:
as a model for a watcher daemon that can launch simd as a subprocess and then read the upgrade log message
to swap binaries as needed. You can pass in information into Plan.Info according to the format
specified here https:
This will allow a properly configured cosmsod daemon to auto-download new binaries and auto-upgrade.
As noted there, this is intended more for full nodes than validators.

Cancelling Upgrades

There are two ways to cancel a planned upgrade - with on-chain governance or off-chain social consensus.
For the first one, there is a CancelSoftwareUpgrade proposal type, which can be voted on and will
remove the scheduled upgrade plan. Of course this requires that the upgrade was known to be a bad idea
well before the upgrade itself, to allow time for a vote. If you want to allow such a possibility, you
should set the upgrade height to be 2 * (votingperiod + depositperiod) + (safety delta) from the beginning of
the first upgrade proposal. Safety delta is the time available from the success of an upgrade proposal
and the realization it was a bad idea (due to external testing). You can also start a CancelSoftwareUpgrade
proposal while the original SoftwareUpgrade proposal is still being voted upon, as long as the voting
period ends after the SoftwareUpgrade proposal.

However, let's assume that we don't realize the upgrade has a bug until shortly before it will occur
(or while we try it out - hitting some panic in the migration). It would seem the blockchain is stuck,
but we need to allow an escape for social consensus to overrule the planned upgrade. To do so, there's
a --unsafe-skip-upgrades flag to the start command, which will cause the node to mark the upgrade
as done upon hitting the planned upgrade height(s), without halting and without actually performing a migration.
If over two-thirds run their nodes with this flag on the old binary, it will allow the chain to continue through
the upgrade with a manual override. (This must be well-documented for anyone syncing from genesis later on).

Example:
	simd start --unsafe-skip-upgrades <height1> <optional_height_2> ... <optional_height_N>

NOTE: Here simd is used as an example binary, replace it with original binary
*/
package upgrade
