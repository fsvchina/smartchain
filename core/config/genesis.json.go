package config

const GenesisJson = `
{
  "genesis_time": "2022-06-29T07:29:56.2166901Z",
  "chain_id": "sc_8888-1",
  "initial_height": "1",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "10000000",
      "time_iota_ms": "30000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000",
      "max_bytes": "1048576"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    },
    "version": {}
  },
  "app_hash": "",
  "app_state": {
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/ethermint.types.v1.EthAccount",
          "base_account": {
            "address": "dex10hrk2aw6ukvvwqglcjm6qrgdru74jpnnhzlq2j",
            "pub_key": null,
            "account_number": "0",
            "sequence": "0"
          },
          "code_hash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
        }
      ]
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "dex10hrk2aw6ukvvwqglcjm6qrgdru74jpnnhzlq2j",
          "coins": [
            {
              "denom": "afsv",
              "amount": "100000000000000000000000000"
            }
          ]
        }
      ],
      "supply": [
        {
          "denom": "afsv",
          "amount": "100000000000000000000000000"
        }
      ],
      "denom_metadata": []
    },
    "capability": {
      "index": "1",
      "owners": []
    },
    "claims": {
      "params": {
        "enable_claims": true,
        "airdrop_start_time": "0001-01-01T00:00:00Z",
        "duration_until_decay": "2629800s",
        "duration_of_decay": "5259600s",
        "claims_denom": "aevmos",
        "authorized_channels": [
          "channel-0",
          "channel-3"
        ],
        "evm_channels": [
          "channel-2"
        ]
      },
      "claims_records": []
    },
    "crisis": {
      "constant_fee": {
        "denom": "afsv",
        "amount": "1000"
      }
    },
    "distribution": {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.010000000000000000",
        "bonus_proposer_reward": "0.040000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": [],
      "previous_proposer": "",
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "epochs": {
      "epochs": [
        {
          "identifier": "week",
          "start_time": "0001-01-01T00:00:00Z",
          "duration": "604800s",
          "current_epoch": "0",
          "current_epoch_start_time": "0001-01-01T00:00:00Z",
          "epoch_counting_started": false,
          "current_epoch_start_height": "0"
        },
        {
          "identifier": "day",
          "start_time": "0001-01-01T00:00:00Z",
          "duration": "86400s",
          "current_epoch": "0",
          "current_epoch_start_time": "0001-01-01T00:00:00Z",
          "epoch_counting_started": false,
          "current_epoch_start_height": "0"
        }
      ]
    },
    "erc20": {
      "params": {
        "enable_erc20": true,
        "enable_evm_hook": true
      },
      "token_pairs": []
    },
    "evidence": {
      "evidence": []
    },
    "evm": {
      "accounts": [],
      "params": {
        "evm_denom": "afsv",
        "enable_create": true,
        "enable_call": true,
        "extra_eips": [],
        "chain_config": {
          "homestead_block": "0",
          "dao_fork_block": "0",
          "dao_fork_support": true,
          "eip150_block": "0",
          "eip150_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
          "eip155_block": "0",
          "eip158_block": "0",
          "byzantium_block": "0",
          "constantinople_block": "0",
          "petersburg_block": "0",
          "istanbul_block": "0",
          "muir_glacier_block": "0",
          "berlin_block": "0",
          "london_block": "0",
          "arrow_glacier_block": "0",
          "merge_fork_block": "0"
        }
      }
    },
    "feegrant": {
      "allowances": []
    },
    "feemarket": {
      "params": {
        "no_base_fee": false,
        "base_fee_change_denominator": 8,
        "elasticity_multiplier": 2,
        "enable_height": "0",
        "base_fee": "1000000000"
      },
      "block_gas": "0"
    },
    "genutil": {
      "gen_txs": [
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "localtestnet",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "0.100000000000000000",
                  "max_rate": "0.200000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "dex10hrk2aw6ukvvwqglcjm6qrgdru74jpnnhzlq2j",
                "validator_address": "dexvaloper10hrk2aw6ukvvwqglcjm6qrgdru74jpnn38vkh0",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "QbFGrzHL4a85Bg3ymGMrMiaUmkqYMrv8Ec9J7RzeqUQ="
                },
                "value": {
                  "denom": "afsv",
                  "amount": "1000000000000000000000"
                }
              }
            ],
            "memo": "78349bc6fec99dafcb9d67d0df82067cd24dfebc@10.5.11.3:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
                  "key": "A+6KUKm3bPb+5hOzmVEUWNM0W6HdX2mYzpOz6MntHmBc"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            }
          },
          "signatures": [
            "zyhmq6VrfCQrGE8woi/BH+xpNJ0VRLYZInSCT9hOAlZ4wiWwiKDyQy8n5i81qfjuPm+PYAu7qcWr1++fCnEOkgE="
          ]
        }
      ]
    },
    "gov": {
      "starting_proposal_id": "1",
      "deposits": [],
      "votes": [],
      "proposals": [],
      "deposit_params": {
        "min_deposit": [
          {
            "denom": "afsv",
            "amount": "10000000"
          }
        ],
        "max_deposit_period": "172800s"
      },
      "voting_params": {
        "voting_period": "172800s"
      },
      "tally_params": {
        "quorum": "0.334000000000000000",
        "threshold": "0.500000000000000000",
        "veto_threshold": "0.334000000000000000"
      }
    },
    "ibc": {
      "client_genesis": {
        "clients": [],
        "clients_consensus": [],
        "clients_metadata": [],
        "params": {
          "allowed_clients": [
            "06-solomachine",
            "07-tendermint"
          ]
        },
        "create_localhost": false,
        "next_client_sequence": "0"
      },
      "connection_genesis": {
        "connections": [],
        "client_connection_paths": [],
        "next_connection_sequence": "0",
        "params": {
          "max_expected_time_per_block": "30000000000"
        }
      },
      "channel_genesis": {
        "channels": [],
        "acknowledgements": [],
        "commitments": [],
        "receipts": [],
        "send_sequences": [],
        "recv_sequences": [],
        "ack_sequences": [],
        "next_channel_sequence": "0"
      }
    },
    "incentives": {
      "params": {
        "enable_incentives": true,
        "allocation_limit": "0.050000000000000000",
        "incentives_epoch_identifier": "week",
        "reward_scaler": "1.200000000000000000"
      },
      "incentives": [],
      "gas_meters": []
    },
    "inflation": {
      "params": {
        "mint_denom": "afsv",
        "exponential_calculation": {
          "a": "300000000.000000000000000000",
          "r": "0.500000000000000000",
          "c": "9375000.000000000000000000",
          "bonding_target": "0.660000000000000000",
          "max_variance": "0.000000000000000000"
        },
        "inflation_distribution": {
          "staking_rewards": "0.533333334000000000",
          "usage_incentives": "0.333333333000000000",
          "community_pool": "0.133333333000000000"
        },
        "enable_inflation": true
      },
      "period": "0",
      "epoch_identifier": "day",
      "epochs_per_period": "365",
      "skipped_epochs": "0"
    },
    "mint": {
      "params": {
        "mint_denom": "afsv"
      }
    },
    "params": null,
    "recovery": {
      "params": {
        "enable_recovery": true,
        "packet_timeout_duration": "14400s"
      }
    },
    "slashing": {
      "params": {
        "signed_blocks_window": "100",
        "min_signed_per_window": "0.500000000000000000",
        "downtime_jail_duration": "600s",
        "slash_fraction_double_sign": "0.050000000000000000",
        "slash_fraction_downtime": "0.010000000000000000"
      },
      "signing_infos": [],
      "missed_blocks": []
    },
    "staking": {
      "params": {
        "unbonding_time": "1814400s",
        "max_validators": 100,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "afsv"
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [],
      "delegations": [],
      "unbonding_delegations": [],
      "redelegations": [],
      "exported": false
    },
    "transfer": {
      "port_id": "transfer",
      "denom_traces": [],
      "params": {
        "send_enabled": true,
        "receive_enabled": true
      }
    },
    "upgrade": {},
    "vesting": {}
  }
}`
