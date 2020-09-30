package audit

import (
	"encoding/json"
	"github.com/ppc64le-cloud/pvsadm/pkg"
	"k8s.io/klog/v2"
	"os"
	"sync"
	"time"
)

var Logger *Audit

func Log(name, op, value string) {
	Logger.Log(name, op, value)
}

type Audit struct {
	file  *os.File
	mutex *sync.Mutex
}

type log struct {
	Name      string    `json:"name"`
	Operation string    `json:"op"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func New(file string) *Audit {
	logfile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		klog.Fatalf("failed to open audit file: %s, err: %v", pkg.Options.AuditFile, err)
	}
	return &Audit{
		file:  logfile,
		mutex: &sync.Mutex{},
	}
}

func (a *Audit) Log(name, op, value string) {
	a.mutex.Lock()
	entry := log{
		Name:      name,
		Operation: op,
		Value:     value,
		Timestamp: time.Now().UTC(),
	}
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		klog.Fatalf("json marshal error: %v", err)
	}
	jsonEntry = append(jsonEntry, '\n')
	if _, err := a.file.Write(jsonEntry); err != nil {
		klog.Fatalf("log failed with error: %v", err)
	}
	a.mutex.Unlock()
}
