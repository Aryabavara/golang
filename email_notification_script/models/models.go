package models

import "time"

type DeploymentStatus struct {
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
