package main
import "fmt"
type camera struct { } 
func (c *camera) Takeapicture() string {
	return "Click"
}
type phone struct { } 
func (p *phone) Call() string {
	return "Ring Ring"
}
type cameraphone struct {
	camera
	phone
}
func main() {
cp:=new(cameraphone)
fmt.Println("Two new features")
fmt.Println("A feature from camera to ",cp.Takeapicture())
fmt.Println("And another from phone to ",cp.Call())
}