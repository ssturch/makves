package storage

type Item struct {
	Id   string `json:"id"`
	Data any    `json:"data"`
}

func (s *Storage) GetInfoByIds(ids []string) ([]Item, error) {
	storeRes, err := s.store.GetInfoByIds(ids)
	if err != nil {
		return nil, err
	}
	res := make([]Item, len(storeRes))
	for i, v := range storeRes {
		res[i] = Item{
			Id:   v.Id,
			Data: v.Data,
		}
	}
	return res, nil
}
