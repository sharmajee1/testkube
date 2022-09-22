package testsources

import (
	testsourcev1 "github.com/kubeshop/testkube-operator/apis/testsource/v1"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MapCRDToAPI maps TestSource CRD to OpenAPI spec TestSource
func MapCRDToAPI(item testsourcev1.TestSource) testkube.TestSource {
	var repository *testkube.Repository
	if item.Spec.Repository != nil {
		repository = &testkube.Repository{
			Type_:  item.Spec.Repository.Type_,
			Uri:    item.Spec.Repository.Uri,
			Branch: item.Spec.Repository.Branch,
			Commit: item.Spec.Repository.Commit,
			Path:   item.Spec.Repository.Path,
		}

		if item.Spec.Repository.UsernameSecret != nil {
			repository.UsernameSecret = &testkube.SecretRef{
				Name: item.Spec.Repository.UsernameSecret.Name,
				Key:  item.Spec.Repository.UsernameSecret.Key,
			}
		}

		if item.Spec.Repository.TokenSecret != nil {
			repository.TokenSecret = &testkube.SecretRef{
				Name: item.Spec.Repository.TokenSecret.Name,
				Key:  item.Spec.Repository.TokenSecret.Key,
			}
		}
	}

	return testkube.TestSource{
		Name:       item.Name,
		Namespace:  item.Namespace,
		Type_:      item.Spec.Type_,
		Uri:        item.Spec.Uri,
		Data:       item.Spec.Data,
		Repository: repository,
		Labels:     item.Labels,
	}
}

// MapAPIToCRD maps OpenAPI spec TestSourceUpsertRequest to CRD TestSource
func MapAPIToCRD(request testkube.TestSourceUpsertRequest) testsourcev1.TestSource {
	var repository *testsourcev1.Repository
	if request.Repository != nil {
		repository = &testsourcev1.Repository{
			Type_:  request.Repository.Type_,
			Uri:    request.Repository.Uri,
			Branch: request.Repository.Branch,
			Commit: request.Repository.Commit,
			Path:   request.Repository.Path,
		}

		if request.Repository.UsernameSecret != nil {
			repository.UsernameSecret = &testsourcev1.SecretRef{
				Name: request.Repository.UsernameSecret.Name,
				Key:  request.Repository.UsernameSecret.Key,
			}
		}

		if request.Repository.TokenSecret != nil {
			repository.TokenSecret = &testsourcev1.SecretRef{
				Name: request.Repository.TokenSecret.Name,
				Key:  request.Repository.TokenSecret.Key,
			}
		}
	}

	return testsourcev1.TestSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name,
			Namespace: request.Namespace,
			Labels:    request.Labels,
		},
		Spec: testsourcev1.TestSourceSpec{
			Type_:      request.Type_,
			Uri:        request.Uri,
			Data:       request.Data,
			Repository: repository,
		},
	}
}
