/*Package contractor(version: "2.0") - Automatically generated by cinp-codegen from /api/v1/Auth/ at 2020-03-08T03:14:55.207366
 */
package contractor

import (
	"reflect"
	cinp "github.com/cinp/go"
)

//AuthUser - Model User(/api/v1/Auth/User)
/*

 */
type AuthUser struct {
	cinp.BaseObject
	cinp *cinp.CInP
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *AuthUser) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
		}
	}
	return &map[string]interface{}{ 
	}
}

// AuthUserNew - Make a new object of Model User
func (service *Contractor) AuthUserNew() *AuthUser {
	return &AuthUser{cinp: service.cinp}
}

// AuthUserNewWithID - Make a new object of Model User
func (service *Contractor) AuthUserNewWithID(id string) *AuthUser {
	result := AuthUser{cinp: service.cinp}
	result.SetID("/api/v1/Auth/User:"+id+":")
	return &result
}

// AuthUserGet - Get function for Model User
func (service *Contractor) AuthUserGet(id string) (*AuthUser, error) {
	object, err := service.cinp.Get("/api/v1/Auth/User:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*AuthUser)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model User
func (object *AuthUser) Create() error {
	if err := object.cinp.Create("/api/v1/Auth/User", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model User
func (object *AuthUser) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model User
func (object *AuthUser) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// AuthUserList - List function for Model User
func (service *Contractor) AuthUserList(filterName string, filterValues map[string]interface{}) <-chan *AuthUser {
	in := service.cinp.ListObjects("/api/v1/Auth/User", reflect.TypeOf(AuthUser{}), filterName, filterValues, 50)
	out := make(chan *AuthUser)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*AuthUser)
		}
	}()
	return out
}

// AuthUserCallLogin calls loginNoneNone
func (service *Contractor) AuthUserCallLogin(username string, password string) (string, error) {
	args := map[string]interface{}{
		"username": username,
		"password": password,
	}
	uri := "/api/v1/Auth/User(login)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// AuthUserCallLogout calls logout
func (service *Contractor) AuthUserCallLogout() (string, error) {
	args := map[string]interface{}{
	}
	uri := "/api/v1/Auth/User(logout)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// AuthUserCallWhoami calls whoami
func (service *Contractor) AuthUserCallWhoami() (string, error) {
	args := map[string]interface{}{
	}
	uri := "/api/v1/Auth/User(whoami)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// AuthUserCallChange_password calls change_passwordNone
func (service *Contractor) AuthUserCallChange_password(password string) (string, error) {
	args := map[string]interface{}{
		"password": password,
	}
	uri := "/api/v1/Auth/User(change_password)"

	result := ""

	if err := service.cinp.Call(uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

func registerAuth(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Auth/User", reflect.TypeOf((*AuthUser)(nil)).Elem())
}