package deploymentresource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
)

var deployTimeoutInMin = 20

func CreateDeployment(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deployment *models.Deployment,
) (*models.Deployment, error) {
	url := baseURL + "deployments?tenant=" + tenant

	res, err := utils.PostRequest(url, *deployment, authHeader, accessToken)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resDeploy := models.Deployment{}
	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return nil, err
	}
	return &resDeploy, err
}

func ReadDeployment(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deploymentName string,
) (*models.Deployment, error) {
	url := baseURL + "deployments/" + deploymentName + "?tenant=" + tenant
	res, err := utils.GetRequest(url, authHeader, accessToken)

	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(resBody))
	}

	if err != nil {
		return nil, err
	}
	resDeploy := models.Deployment{}

	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return nil, err
	}
	return &resDeploy, nil
}

func UpdateDeployment(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deployment *models.Deployment,
) (*models.Deployment, error) {
	url := baseURL + "deployments/" + *deployment.Name + "?tenant=" + tenant

	res, err := utils.PutRequest(url, *deployment, authHeader, accessToken)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 204 && res.StatusCode != 200 {
		return nil, errors.New(string(resBody))
	}

	resDeploy := models.Deployment{}
	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return nil, err
	}

	return &resDeploy, nil
}

func DeleteDeployment(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deploymentName string,
) error {
	url := baseURL + "deployments/" + deploymentName + "?tenant=" + tenant

	err := utils.DeleteRequest(url, authHeader, accessToken)
	if err != nil {
		return err
	}
	return nil
}

func GetDeployment(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deploymentName string,
) (*models.Deployment, error) {
	url := baseURL + "deployments/" + deploymentName + "?tenant=" + tenant
	res, err := utils.GetRequest(url, authHeader, accessToken)

	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(resBody))
	}

	if err != nil {
		return nil, err
	}

	resDeploy := models.Deployment{}
	err = json.Unmarshal(resBody, &resDeploy)
	if err != nil {
		return nil, err
	}

	return &resDeploy, nil
}

func CheckDeployment(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deploymentName string,
	action string,
) (*models.Deployment, error) {
	for i := 0; i < deployTimeoutInMin*60; i += 10 {
		time.Sleep(10 * time.Second)
		deployment, err := ReadDeployment(ctx, baseURL, tenant, authHeader, accessToken, deploymentName)
		if err != nil {
			continue
		} else {
			statusID := deployment.StatusID
			if statusID == utils.STATUS_RUNNING_ID || statusID == utils.STATUS_FAILED_ID {
				return deployment, nil
			} else if statusID == utils.STATUS_FAILED_ID {
				return nil, fmt.Errorf("%s failed", action)
			}
		}
	}
	return nil, fmt.Errorf("%s timed out", action)
}

func CheckDeletion(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	deploymentName string,
) error {
	for i := 0; i < deployTimeoutInMin*60; i += 10 {
		time.Sleep(10 * time.Second)
		deployment, err := ReadDeployment(ctx, baseURL, tenant, authHeader, accessToken, deploymentName)
		if err != nil {
			msg := swaggerRes{}
			err = json.Unmarshal([]byte(err.Error()), &msg)
			if err == nil && msg.Code == 404 {
				return nil
			}
		} else {
			statusID := deployment.StatusID
			if statusID == utils.STATUS_DELETED_ID {
				return nil
			} else if statusID == utils.STATUS_FAILED_ID {
				return errors.New("Deletion failed")
			}
		}
	}
	return errors.New("Deletion timed out")
}

type swaggerRes struct {
	Code    int
	Message string
}
