package models

type Entity struct {
	ID                        int      `csv:"id"`
	UID                       string   `csv:"uid"`
	Domain                    string   `csv:"domain"`
	CN                        string   `csv:"cn"`
	Department                string   `csv:"department"`
	Title                     string   `csv:"title"`
	Who                       string   `csv:"who"`
	LogonCount                int      `csv:"logon_count"`
	NumLogons7                int      `csv:"num_logons7"`
	NumShare7                 int      `csv:"num_share7"`
	NumFile7                  int      `csv:"num_file7"`
	NumAd7                    int      `csv:"num_ad7"`
	NumN7                     int      `csv:"num_n7"`
	NumLogons14               int      `csv:"num_logons14"`
	NumShare14                int      `csv:"num_share14"`
	NumFile14                 int      `csv:"num_file14"`
	NumAd14                   int      `csv:"num_ad14"`
	NumN14                    int      `csv:"num_n14"`
	NumLogons30               int      `csv:"num_logons30"`
	NumShare30                int      `csv:"num_share30"`
	NumFile30                 int      `csv:"num_file30"`
	NumAd30                   int      `csv:"num_ad30"`
	NumN30                    int      `csv:"num_n30"`
	NumLogons150              int      `csv:"num_logons150"`
	NumShare150               int      `csv:"num_share150"`
	NumFile150                int      `csv:"num_file150"`
	NumAd150                  int      `csv:"num_ad150"`
	NumN150                   int      `csv:"num_n150"`
	NumLogons365              int      `csv:"num_logons365"`
	NumShare365               int      `csv:"num_share365"`
	NumFile365                int      `csv:"num_file365"`
	NumAd365                  int      `csv:"num_ad365"`
	NumN365                   int      `csv:"num_n365"`
	HasUserPrincipalName      bool     `csv:"has_user_principal_name"`
	HasMail                   bool     `csv:"has_mail"`
	HasPhone                  bool     `csv:"has_phone"`
	FlagDisabled              bool     `csv:"flag_disabled"`
	FlagLockout               bool     `csv:"flag_lockout"`
	FlagPasswordNotRequired   bool     `csv:"flag_password_not_required"`
	FlagPasswordCantChange    bool     `csv:"flag_password_cant_change"`
	FlagDontExpirePassword    bool     `csv:"flag_dont_expire_password"`
	OwnedFiles                int      `csv:"owned_files"`
	NumMailboxes              int      `csv:"num_mailboxes"`
	NumMemberOfGroups         int      `csv:"num_member_of_groups"`
	NumMemberOfIndirectGroups int      `csv:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIDs []string `csv:"member_of_indirect_groups_i_ds"`
	MemberOfGroupsIDs         []string `csv:"member_of_groups_i_ds"`
	IsAdmin                   bool     `csv:"is_admin"`
	IsService                 bool     `csv:"is_service"`
}
