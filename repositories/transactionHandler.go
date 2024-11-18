package repositories

func CommitTransaction(user *UsersRepository) error {
	transaction := user.transaction
	user.transaction = nil

	return transaction.Commit()
}
