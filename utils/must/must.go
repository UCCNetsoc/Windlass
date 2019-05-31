package must

import (
	"fmt"
)

func Do(funs ...func() error) {
	for _, fun := range funs {
		do(fun)
	}
}

func do(fun func() error) {
	if err := fun(); err != nil {
		panic(fmt.Sprintf("failed to start app: %v", err))
	}
}
