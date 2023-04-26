package needle

import "github.com/IBAX-io/needle/vm"

func New() *vm.VM {
	return vm.GetVM()
}
