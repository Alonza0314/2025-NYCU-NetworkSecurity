package main_test

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"golang.org/x/crypto/sha3"
)

var md5Cases = []struct {
	hash     string
	password string
}{
	{
		hash:     "e49201c3a8f548902b9ae9f16638f879",
		password: "0890003871",
	},
	{
		hash:     "19cf9dda4107b300d3218702df95c76d",
		password: "nailz07",
	},
	{
		hash:     "c6281df39e8ade06c6cc9e0095fd5c0f",
		password: "rksmbffs",
	},
	{
		hash:     "a54034981409ed58d584dc9051853ddb",
		password: "hidalgo212",
	},
	{
		hash:     "f58291f81868320f11235d9b9d416115",
		password: "aq12wsxz",
	},
	{
		hash:     "ce1c96461fbb2ad92fffcafafe85d0d1",
		password: "CAROLIAN",
	},
	{
		hash:     "c6177167ebb2c37352c3a63f6fa0c39d",
		password: "19821983",
	},
	{
		hash:     "5993428babd2cb253834e06de1800916",
		password: "netopia",
	},
	{
		hash:     "bebc51b6f0bbd5da67950200a89026f6",
		password: "Autumn2018",
	},
	{
		hash:     "456c5a41af2eb09ac0ba0eb64f614887",
		password: "eeyore",
	},
}

var sha3Cases = []struct {
	hash     string
	password string
}{
	{
		hash:     "1074f17769cc2dfc0d65f713a7d8c4fd97fc78c69cfa13263b07b0e40b3cf83a",
		password: "sweetlove",
	},
	{
		hash:     "94f72dc2ea6bfae657b0ee3d5adb992aa669f6c4141717344e24e873dc09be04",
		password: "shunkoko",
	},
	{
		hash:     "19c743dc300d52fc93b5ee8c6d224f3beb8a05079e6439855cdae7e55bf16ef0",
		password: "mrzdale08",
	},
	{
		hash:     "20e5b0556c431db9a147c3f73a0ae03d12f5ef391d277cd59ff0f2dd98198ec5",
		password: "minot24601",
	},
	{
		hash:     "a44cf105063b06bbb160c22058e9c3137c8ef424ae72f981d73b10fdc743026f",
		password: "loverboydj242",
	},
	{
		hash:     "74151544815c4a0153c2e7dfabcfd066d510d6996148d6c02f246c9c497bd15c",
		password: "bear1194",
	},
	// {
	// 	hash:     "745af7302284f80ddadf6893f64e247334aa899bfe90512a59aa41ea2863f56a",
	// 	password: "",
	// },
	// {
	// 	hash:     "9d34ebe967a790ada61cfa2b4e16671bfb18f0ff59296f24a0eec20dacc5ece3",
	// 	password: "",
	// },
	{
		hash:     "0ecd9ac47c8e4b059c2b97db9657f80f203454ac8fcb01976e1decdb30af2510",
		password: "bambam",
	},
	{
		hash:     "3b2918324171f88304baee77d71cc0abd40e12f16f9a22404736000f00a7c7b6",
		password: "maddie1",
	},
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func sha3Hash(text string) string {
	hash := sha3.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func TestMain(t *testing.T) {
	t.Run("MD5", func(t *testing.T) {
		for _, test := range md5Cases {
			if md5Hash(test.password) != test.hash {
				t.Errorf("Expected %s, got %s", test.hash, md5Hash(test.password))
			}
		}
	})

	t.Run("SHA3-256", func(t *testing.T) {
		for _, test := range sha3Cases {
			if sha3Hash(test.password) != test.hash {
				t.Errorf("Expected %s, got %s", test.hash, sha3Hash(test.password))
			}
		}
	})
}
