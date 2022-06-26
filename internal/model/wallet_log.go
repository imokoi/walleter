package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type ERC20TokenCollection struct {
	Items []ERC20TokenData `json:"items"`
}

func (item ERC20TokenCollection) Value() (driver.Value, error) {
	b, err := json.Marshal(item)
	return string(b), err
}

func (item *ERC20TokenCollection) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), item)
}

type ERC20TokenData struct {
	TokenType string  `json:"token_type" gorm:"type:varchar(20)"`
	Amount    float64 `json:"amount"`
	Decimal   uint    `json:"decimal"`
}

func (item ERC20TokenData) Value() (driver.Value, error) {
	b, err := json.Marshal(item)
	return string(b), err
}

func (item *ERC20TokenData) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), item)
}

// ERC20WalletLog Wallet flow log
type ERC20WalletLog struct {
	gorm.Model     `swagger-ignore:"true"`
	AccountId      uint                 `json:"account_id"`
	BusinessModule string               `json:"business_module" gorm:"type:varchar(64);not null;comment:'业务模块'"`
	ActionType     string               `json:"action_type" gorm:"type:varchar(64);not null;comment:'操作类型'"`
	Tokens         ERC20TokenCollection `json:"tokens" gorm:"type:json;not null"`
	Fees           ERC20TokenCollection `json:"fees" gorm:"type:json;"`
	Status         string               `json:"status" gorm:"type:varchar(64);not null;comment:处理状态"`
	OriginalWallet Wallet               `json:"original_wallet" gorm:"type:json;not null;comment:'变更前的钱包数据'"`
	SettledWallet  Wallet               `json:"settled_wallet" gorm:"type:json;not null;comment:'变更后的钱包数据'"`
}

// ERC1155WalletLog Wallet flow log
type ERC1155WalletLog struct {
	gorm.Model     `swagger-ignore:"true"`
	AccountId      uint                 `json:"account_id"`
	BusinessModule string               `json:"business_module" gorm:"type:varchar(64);not null;comment:'业务模块'"`
	ActionType     string               `json:"action_type" gorm:"type:varchar(64);not null;comment:'操作类型'"`
	Ids            string               `json:"ids"`
	Values         string               `json:"values"`
	Fees           ERC20TokenCollection `json:"fees" gorm:"type:json;"`
	Status         string               `json:"status" gorm:"type:varchar(10);not null;comment:处理状态"`
	OriginalWallet Wallet               `json:"original_wallet" gorm:"type:json;not null;comment:'变更前的钱包数据'"`
	SettledWallet  Wallet               `json:"settled_wallet" gorm:"type:json;comment:'变更后的钱包数据'"`
}

type erc20WalletLogDAO struct{}

var ERC20WalletLogDAO = &erc20WalletLogDAO{}

func (s erc20WalletLogDAO) InsertERC20WalletLog(db *gorm.DB, erc20Log ERC20WalletLog) (ERC20WalletLog, error) {
	err := db.Create(&erc20Log).Error
	if err != nil {
		return ERC20WalletLog{}, err
	}
	return erc20Log, nil
}

func (s erc20WalletLogDAO) UpdateERC20WalletLogStatus(db *gorm.DB, newLog ERC20WalletLog) (ERC20WalletLog, error) {
	err := db.Save(&newLog).Error
	if err != nil {
		return ERC20WalletLog{}, err
	}
	return newLog, nil
}

type erc1155WalletLogDAO struct{}

var ERC1155WalletLogDAO = &erc1155WalletLogDAO{}

func (s erc1155WalletLogDAO) InsertERC1155WalletLog(db *gorm.DB, erc1155Log ERC1155WalletLog) (ERC1155WalletLog, error) {
	err := db.Create(&erc1155Log).Error
	if err != nil {
		return ERC1155WalletLog{}, err
	}
	return erc1155Log, nil
}

func (s erc1155WalletLogDAO) UpdateERC1155WalletLogStatus(db *gorm.DB, newLog ERC1155WalletLog) (ERC1155WalletLog, error) {
	err := db.Save(&newLog).Error
	if err != nil {
		return ERC1155WalletLog{}, err
	}
	return newLog, nil
}
