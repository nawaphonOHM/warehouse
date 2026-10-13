package transaction

func Transaction(operation func()) bool {

	operation()

	return true
}
