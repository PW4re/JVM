package runtime

type LocalVariables struct {
	variables []Value
}

func initLocalVariables(this *Object, vars []Value) LocalVariables {
	return LocalVariables{
		variables: append([]Value{this}, vars),
	}
}
