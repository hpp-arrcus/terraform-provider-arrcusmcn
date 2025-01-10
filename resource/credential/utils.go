package credentialresource

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
)

func CreateCred(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	cred *models.Credentials,
) (*models.Credentials, error) {
	url := baseURL + "cloud_credentials?tenant=" + tenant

	var err error

	res, err := utils.PostRequest(url, *cred, authHeader, accessToken)

	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resCred := models.Credentials{}
	err = json.Unmarshal(resBody, &resCred)
	if err != nil {
		return nil, err
	}
	return &resCred, nil
}

func GetCred(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	credName string,
) (*models.Credentials, error) {
	url := baseURL + "cloud_credentials/" + credName + "?tenant=" + tenant

	res, err := utils.GetRequest(url, authHeader, accessToken)
	// return diag.FromErr(err)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resCred := models.Credentials{}
	err = json.Unmarshal(resBody, &resCred)
	if err != nil {
		return nil, err
	}

	return &resCred, nil
}

func UpdateCred(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	cred *models.Credentials,
) (*models.Credentials, error) {
	url := baseURL + "cloud_credentials/" + *cred.Name + "?tenant=" + tenant

	var err error

	res, err := utils.PutRequest(url, *cred, authHeader, accessToken)

	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resCred := models.Credentials{}
	err = json.Unmarshal(resBody, &resCred)
	if err != nil {
		return nil, err
	}
	return &resCred, nil
}

func DeleteCred(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	credName string,
) error {
	url := baseURL + "cloud_credentials/" + credName + "?tenant=" + tenant

	err := utils.DeleteRequest(url, authHeader, accessToken)
	if err != nil {
		return err
	}
	return nil
}
