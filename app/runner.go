package app

type Runner interface {
	Service
	Run()
}
