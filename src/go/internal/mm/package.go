package mm

import "context"

func ReadScriptFromFile(filename string) error {
	return DefaultMM.ReadScriptFromFile(filename)
}

func ClearNamespace(ns string) error {
	return DefaultMM.ClearNamespace(ns)
}

func LaunchVMs(ns string) error {
	return DefaultMM.LaunchVMs(ns)
}

func GetLaunchProgress(ns string, expected int) (float64, error) {
	return DefaultMM.GetLaunchProgress(ns, expected)
}

func GetVMInfo(opts ...Option) VMs {
	return DefaultMM.GetVMInfo(opts...)
}

func GetVMScreenshot(opts ...Option) ([]byte, error) {
	return DefaultMM.GetVMScreenshot(opts...)
}

func GetVNCEndpoint(opts ...Option) (string, error) {
	return DefaultMM.GetVNCEndpoint(opts...)
}

func StartVM(opts ...Option) error {
	return DefaultMM.StartVM(opts...)
}

func StopVM(opts ...Option) error {
	return DefaultMM.StopVM(opts...)
}

func RedeployVM(opts ...Option) error {
	return DefaultMM.RedeployVM(opts...)
}

func KillVM(opts ...Option) error {
	return DefaultMM.KillVM(opts...)
}

func GetVMHost(opts ...Option) (string, error) {
	return DefaultMM.GetVMHost(opts...)
}

func ConnectVMInterface(opts ...Option) error {
	return DefaultMM.ConnectVMInterface(opts...)
}

func DisconnectVMInterface(opts ...Option) error {
	return DefaultMM.DisconnectVMInterface(opts...)
}

func StartVMCapture(opts ...Option) error {
	return DefaultMM.StartVMCapture(opts...)
}

func StopVMCapture(opts ...Option) error {
	return DefaultMM.StopVMCapture(opts...)
}

func GetExperimentCaptures(opts ...Option) []Capture {
	return DefaultMM.GetExperimentCaptures(opts...)
}

func GetVMCaptures(opts ...Option) []Capture {
	return DefaultMM.GetVMCaptures(opts...)
}

func GetClusterHosts(schedOnly bool) (Hosts, error) {
	return DefaultMM.GetClusterHosts(schedOnly)
}

func Headnode() string {
	return DefaultMM.Headnode()
}

func IsHeadnode(node string) bool {
	return DefaultMM.IsHeadnode(node)
}

func GetVLANs(opts ...Option) (map[string]int, error) {
	return DefaultMM.GetVLANs(opts...)
}

func IsC2ClientActive(opts ...C2Option) error {
	return DefaultMM.IsC2ClientActive(opts...)
}

func ExecC2Command(opts ...C2Option) (string, error) {
	return DefaultMM.ExecC2Command(opts...)
}

func WaitForC2Response(ctx context.Context, opts ...C2Option) (string, error) {
	return DefaultMM.WaitForC2Response(ctx, opts...)
}

func ClearC2Responses(opts ...C2Option) error {
	return DefaultMM.ClearC2Responses(opts...)
}
