package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Metadata struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Metadata Metadata `json:"metadata"`
	Data     string   `json:"response"`
}

const consid = ""    //Enter consID
const timestamp = "" //Enter timeStamp
const signature = "" //Enter signature
const key = ""       //Enter userKey

func main() {
	router := gin.Default()

	router.GET("/ref/poli", getPoli)
	router.GET("/ref/dokter", getDokter)
	router.GET("/jadwaldokter/kodepoli/:kodepoli/tanggal/:tanggal", getJadwalDokter)

	router.Run(":8080")
}

func getPoli(c *gin.Context) {
	url := "https://apijkn-dev.bpjs-kesehatan.go.id/antreanrs_dev/ref/poli"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
		return
	}

	req.Header.Set("x-cons-id", consid)
	req.Header.Set("x-timestamp", timestamp)
	req.Header.Set("x-signature", signature)
	req.Header.Set("user_key", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to make request"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unexpected response status code: %d", resp.StatusCode)})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response body"})
		return
	}

	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode response body"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func getDokter(c *gin.Context) {
	url := "https://apijkn-dev.bpjs-kesehatan.go.id/antreanrs_dev/ref/dokter"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
		return
	}

	req.Header.Set("x-cons-id", consid)
	req.Header.Set("x-timestamp", timestamp)
	req.Header.Set("x-signature", signature)
	req.Header.Set("user_key", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to make request"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unexpected response status code: %d", resp.StatusCode)})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response body"})
		return
	}

	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode response body"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func getJadwalDokter(c *gin.Context) {
	kodepoli := c.Param("kodepoli")
	tanggal := c.Param("tanggal")

	url := "https://apijkn-dev.bpjs-kesehatan.go.id/antreanrs_dev/jadwaldokter/kodepoli/" + kodepoli + "/tanggal/" + tanggal
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
		return
	}

	req.Header.Set("x-cons-id", consid)
	req.Header.Set("x-timestamp", timestamp)
	req.Header.Set("x-signature", signature)
	req.Header.Set("user_key", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to make request"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unexpected response status code: %d", resp.StatusCode)})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response body"})
		return
	}

	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode response body"})
		return
	}

	c.JSON(http.StatusOK, response)
}
