package proc

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// TODO: add https://github.com/jbenet/goprocess
type ProcessManager struct {
	exePath     string
	args        []string
	cmd         *exec.Cmd
	logFileName string
	logFile     *os.File
}

func NewProcessManager(exePath string, args []string, lfn string) *ProcessManager {
	return &ProcessManager{
		exePath:     exePath,
		args:        args,
		logFileName: lfn,
	}
}

func (pm *ProcessManager) Start() error {
	pm.cmd = exec.Command(pm.exePath, pm.args...)

	// TODO: what about Init() logs?
	lf, err := os.OpenFile(pm.logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	pm.cmd.Stdout = lf
	pm.cmd.Stderr = lf
	pm.logFile = lf

	err = pm.cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (pm *ProcessManager) Stop() error {
	if pm.cmd == nil {
		return fmt.Errorf("the process is not running")
	}

	if err := pm.logFile.Close(); err != nil {
		return err
	}
	return pm.cmd.Process.Signal(os.Interrupt)
}

// TODO: support windows
func (pm *ProcessManager) pause() error {
	if pm.cmd == nil || pm.cmd.Process == nil {
		return fmt.Errorf("Process not running")
	}
	err := pm.cmd.Process.Signal(syscall.SIGSTOP)
	if err != nil {
		return err
	}
	return nil
}

func (pm *ProcessManager) resume() error {
	if pm.cmd == nil || pm.cmd.Process == nil {
		return fmt.Errorf("Process not running")
	}
	err := pm.cmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}
	return nil
}
