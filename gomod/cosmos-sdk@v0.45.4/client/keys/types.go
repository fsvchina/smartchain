package keys




type AddNewKey struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
	Account  int    `json:"account,string,omitempty"`
	Index    int    `json:"index,string,omitempty"`
}


func NewAddNewKey(name, password, mnemonic string, account, index int) AddNewKey {
	return AddNewKey{
		Name:     name,
		Password: password,
		Mnemonic: mnemonic,
		Account:  account,
		Index:    index,
	}
}


type RecoverKey struct {
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
	Account  int    `json:"account,string,omitempty"`
	Index    int    `json:"index,string,omitempty"`
}


func NewRecoverKey(password, mnemonic string, account, index int) RecoverKey {
	return RecoverKey{Password: password, Mnemonic: mnemonic, Account: account, Index: index}
}


type UpdateKeyReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}


func NewUpdateKeyReq(old, new string) UpdateKeyReq {
	return UpdateKeyReq{OldPassword: old, NewPassword: new}
}


type DeleteKeyReq struct {
	Password string `json:"password"`
}


func NewDeleteKeyReq(password string) DeleteKeyReq { return DeleteKeyReq{Password: password} }
