package main

import (
	"encoding/json"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	b64 "encoding/base64"
)

type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

type Token struct {
	Id                 *big.Int       `json:"id"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	ImageEncodedString string         `json:"image"`
	ImageSVG           string         `json:"image_svg"`
	Attributes         []Attribute    `json:"attributes"`
	Place              IPlacesPlace   `json:"place"`
	Owner              common.Address `json:"owner"`
}

func ethInstance() *Places {
	rpc := os.Getenv("ETH_NODE")
	if len(rpc) == 0 {
		log.Fatal("ETH_NODE env variable not set. It should look something like: https://mainnet.infura.io/v3/YOUR_API_KEY")
		os.Exit(1)
	}
	client, err := ethclient.Dial(rpc)
	Ok(err)

	// Places Mainnet address
	address := common.HexToAddress("0xC9CA129DC3a299aF68A215d85771630aec4C3C2b")
	instance, err := NewPlaces(address, client)
	if err != nil {
		log.Fatal("Cannot connect to contract. Is your RPC Node set to mainnet?")
		os.Exit(1)
	}

	return instance
}

func getToken(tokenIdString string) *Token {
	instance := ethInstance()

	tokenIdInt, err := strconv.ParseInt(tokenIdString, 0, 64)
	Ok(err)

	tokenId := big.NewInt(tokenIdInt)

	place, err := instance.GetPlace(nil, tokenId)
	Ok(err)

	owner, err := instance.OwnerOf(nil, tokenId)
	Ok(err)

	token := instance.parseTokenURI(tokenId)
	token.Id = tokenId
	token.Place = place
	token.Owner = owner

	return token
}

func (instance *Places) parseTokenURI(tokenId *big.Int) *Token {
	tokenBase64, err := instance.TokenURI(nil, tokenId)
	Ok(err)

	tokenBase64Split := strings.Split(tokenBase64, ",")
	tokenDecoded, _ := b64.StdEncoding.DecodeString(tokenBase64Split[1])
	tokenDecordedString := string(tokenDecoded)

	var token Token
	if err := json.Unmarshal([]byte(tokenDecordedString), &token); err != nil {
		panic(err)
	}

	imageString := strings.Split(token.ImageEncodedString, ",")
	imageDecoded, _ := b64.StdEncoding.DecodeString(imageString[1])
	token.ImageSVG = string(imageDecoded)

	return &token
}

func (token *Token) displayToken() {
	f, err := os.CreateTemp("", "place.*.svg")
	Ok(err)

	defer os.Remove(f.Name())
	f.Write([]byte(token.ImageSVG))
	f.Close()

	_, err = exec.Command("/usr/bin/open", f.Name()).Output()
	// println(string(output))
	Ok(err)

	// Sleeping so that the file can be read
	time.Sleep(time.Second * 1)
}

func (token *Token) createJson() {
	bytes, err := json.MarshalIndent(token, "", "\t")
	Ok(err)

	println(string(bytes))
}
