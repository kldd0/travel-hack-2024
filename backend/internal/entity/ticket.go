package entity

type Ticket struct { //сама заявка
	TicketId     string       `json:"ticket_id" db:"ticket_id"`         //номер заявки
	Count        int          `json:"count" db:"count"`                 //кол-во людей в заявке
	GroupDates   TourDate     `json:"group_dates" db:"group_dates"`     //выбираем дату
	TicketHolder TicketHolder `json:"ticket_holder" db:"ticket_holder"` //инфа о заявщике
	MatesInfo    []Mate       `json:"mates_info" db:"mates_info"`
}

type TicketHolder struct { // инфа о том, на кого оформлена заявка
	HolderSurname string // фамилия заявщика
	HolderName    string // имя заявщика
	Age           int    //возраст
	PhoneNumber   string //номер телефона
	Email         string //почта
	WantEvents    bool   //хочет ли рассылку
}

type Mate struct {
	MateSurname string // фамилия побочного участника
	MateName    string // имя побочного участника
	Age         int    //возраст
}
