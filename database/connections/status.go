package connections

func GetCurrentConnectionStatus() ConnectionStatus {
	return currentConnectionStatus
}

func SetConnectionStatus(status ConnectionStatus) {
	currentConnectionStatus = status
}
