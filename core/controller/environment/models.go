package environment

import "g.hz.netease.com/horizon/pkg/environment/models"

type Environment struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Environments []*Environment

// ofEnvironmentModels []*models.Environment to []*Environment
func ofEnvironmentModels(envs []*models.Environment) Environments {
	environments := make(Environments, 0)
	for _, env := range envs {
		environments = append(environments, &Environment{
			Name:        env.Name,
			DisplayName: env.DisplayName,
		})
	}
	return environments
}
