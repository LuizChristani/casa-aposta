package connections

// Função para obter o status atual da conexão
func GetCurrentConnectionStatus() ConnectionStatus {
	return currentConnectionStatus
}

// Função para definir o status da conexão
func SetConnectionStatus(status ConnectionStatus) {
	currentConnectionStatus = status
}