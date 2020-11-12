package main_test

import "testing"

func TestSum(t *testing.T) {
	t.Log("executing", "TestSum(1)")
	t.Log("executing", "TestSum(2)")
}

func TestMin(t *testing.T) {
	t.Log("executing", "TestMin")
}

// return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	if file, e := os.OpenFile("terraform-provider-dynatrace.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); e != nil {
// 		panic(errors.New("Failed to open log file"))
// 	} else {
// 		log.SetOutput(file)
// 	}
// 	return fn(ctx, d, m)
// }
