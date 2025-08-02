package znakdb

// здесь пинг вызывается в сессии и он не закрывает ее
// и по алгоритму, во всех методах пакета надо закрывать сессию обязательно!
func (z *DbZnak) Example() (err error) {
	sess := z.dbSession
	defer sess.Close()

	return sess.Ping()
}
