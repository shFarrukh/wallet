package wallet

import "testing"

func TestService_FindAccoundById_Method_NotFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(3)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindPaymentByID_success(t *testing.T) {
	//создаем сервис
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_Reject_fail(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = svc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := svc.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	editPayID := pay.ID + "l"
	err = svc.Reject(editPayID)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_Reject_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = svc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := svc.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	editPayID := pay.ID 
	err = svc.Reject(editPayID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}
func TestService_Repeat_success_user(t *testing.T) {
	//создаем сервис
	s := newTestServiceUser()
	s.RegisterAccount("+9922000000")
	account, err :=s.FindAccountByID(1)
	if err != nil {
	t.Error(err)
	return
	}
	//пополняем баланс
	err = s.Deposit(account.ID, 1000_00)
	if err != nil {
	t.Errorf("\ngot > %v \nwant > nil", err)
	}
	//pay
	payment, err := s.Pay(account.ID, 100_00, "auto")
	if err != nil {
	t.Errorf("\ngot > %v \nwant > nil", err)
	}
	
	pay, err := s.FindPaymentByID(payment.ID)
	if err != nil {
	t.Errorf("\ngot > %v \nwant > nil", err)
	}
	
	pay, err = s.Repeat(pay.ID)
	if err != nil {
	t.Errorf("Repeat(): can't payment for an account(%v), error(%v)",pay.ID, err)
	}
	}
	func TestService_FavoritePayment_success_user(t *testing.T) {
		//создаем сервис
		var s Service
	
		account, err := s.RegisterAccount("+9922000000")
		if err != nil {
			t.Errorf("method RegisterAccount return not nil error, account=>%v", account)
			return
		}
		//пополняем баланс
		err = s.Deposit(account.ID, 1000_00)
		if err != nil {
			t.Errorf("method Deposit return not nil error, error=>%v", err)
		}
		//pay
		payment, err := s.Pay(account.ID, 100_00, "auto")
		if err != nil {
			t.Errorf("method Pay return not nil error, account=>%v", account)
		}
		//edit favorite
		favorite, err := s.FavoritePayment(payment.ID, "auto")
		if err != nil {
			t.Errorf("method FavoritePayment returned not nil error, favorite=>%v", favorite)
		}
	
		paymentFavorite, err := s.PayFromFavorite(favorite.ID)
		if err != nil {
			t.Errorf("method PayFromFavorite returned nil, paymentFavorite(%v)", paymentFavorite)
		}
	}