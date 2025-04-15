package dom

type FuncI interface {
	// Release frees up resources allocated for the function.
	// The function must not be invoked after calling Release.
	// It is allowed to call Release while the function is still running.
	Release()
}
