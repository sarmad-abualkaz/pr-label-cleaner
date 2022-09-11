package gh

import (
    "context"
    "fmt"
    "net/http"
    "testing"

)

func TestRemoveLable(t *testing.T){

    type endpoint struct{
		label    string
        owner    string
        prNumber int
        repo     string        
    }
    
    type args struct{
		label    string
        owner    string
        prNumber int
        repo     string
	}
	
	tests := []struct {
		name       string
		args       args
        endpoint   endpoint
        wantedErr  error
	}{
		{
            name: "should succefully remove label",
            args: args{
                label:    "test",
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            endpoint: endpoint{
                label:    "test",
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            wantedErr:  nil,
        },
		{
            name: "should fail to find label",
            args: args{
                label:    "test",
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            endpoint: endpoint{
                label:    "test2",
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            wantedErr:  fmt.Errorf("label not found"),
        },
    }

    for _, tt := range tests {
        
        t.Run(tt.name, func(t *testing.T) {

            client, mux, _, teardown := setup()
            defer teardown()

            ctx := context.Background()
                
            apiurl := fmt.Sprintf("/repos/%s/%s/issues/%d/labels/%s", tt.endpoint.owner, tt.endpoint.repo, tt.endpoint.prNumber, tt.endpoint.label)
            
            mux.HandleFunc(apiurl, func(w http.ResponseWriter, r *http.Request) {
                testMethod(t, r, "DELETE")
            })
            
            err := RemoveLable(ctx, client, tt.args.label, tt.args.owner, tt.args.prNumber, tt.args.repo)

            if tt.wantedErr == nil {
                
                if err != nil {
                    t.Errorf("RemoveLable returned error: %v expect %v", err, tt.wantedErr)
                }

            } else if err.Error() != tt.wantedErr.Error() {
                
                t.Errorf("RemoveLable returned error: %v expect %v", err, tt.wantedErr)
            
            }

        })
    }
}
