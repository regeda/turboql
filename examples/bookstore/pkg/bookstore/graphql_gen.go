package bookstore

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/regeda/turboql/pkg/batcher"
	"github.com/regeda/turboql/pkg/graphqlx/scalar"
	"github.com/stephenafamo/scan/pgxscan"
)

type Address struct {
	Address_id    int
	Street_number string
	Street_name   string
	City          string
	Country_id    int
}
type Address_status struct {
	Status_id      int
	Address_status string
}
type Author struct {
	Author_id   int
	Author_name string
}
type Book struct {
	Book_id          int
	Title            string
	Isbn13           string
	Language_id      int
	Num_pages        int
	Publication_date time.Time
	Publisher_id     int
}
type Book_author struct {
	Book_id   int
	Author_id int
}
type Book_language struct {
	Language_id   int
	Language_code string
	Language_name string
}
type Country struct {
	Country_id   int
	Country_name string
}
type Cust_order struct {
	Order_id           int
	Order_date         time.Time
	Customer_id        int
	Shipping_method_id int
	Dest_address_id    int
}
type Customer struct {
	Customer_id int
	First_name  string
	Last_name   string
	Email       string
}
type Customer_address struct {
	Customer_id int
	Address_id  int
	Status_id   int
}
type Order_history struct {
	History_id  int
	Order_id    int
	Status_id   int
	Status_date time.Time
}
type Order_line struct {
	Line_id  int
	Order_id int
	Book_id  int
	Price    pgtype.Numeric
}
type Order_status struct {
	Status_id    int
	Status_value string
}
type Publisher struct {
	Publisher_id   int
	Publisher_name string
}
type Shipping_method struct {
	Method_id   int
	Method_name string
	Cost        pgtype.Numeric
}

func CreateFields(pq pgxscan.Queryer) graphql.Fields {
	addressType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"address_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).Address_id, nil
				},
			},
			"street_number": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).Street_number, nil
				},
			},
			"street_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).Street_name, nil
				},
			},
			"city": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).City, nil
				},
			},
			"country_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).Country_id, nil
				},
			},
		},
	})
	address_statusType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Address_status",
		Fields: graphql.Fields{
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address_status).Status_id, nil
				},
			},
			"address_status": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address_status).Address_status, nil
				},
			},
		},
	})
	authorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"author_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Author).Author_id, nil
				},
			},
			"author_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Author).Author_name, nil
				},
			},
		},
	})
	bookType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"book_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Book_id, nil
				},
			},
			"title": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Title, nil
				},
			},
			"isbn13": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Isbn13, nil
				},
			},
			"language_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Language_id, nil
				},
			},
			"num_pages": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Num_pages, nil
				},
			},
			"publication_date": &graphql.Field{
				Type: scalar.Date,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Publication_date, nil
				},
			},
			"publisher_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).Publisher_id, nil
				},
			},
		},
	})
	book_authorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book_author",
		Fields: graphql.Fields{
			"book_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book_author).Book_id, nil
				},
			},
			"author_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book_author).Author_id, nil
				},
			},
		},
	})
	book_languageType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book_language",
		Fields: graphql.Fields{
			"language_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book_language).Language_id, nil
				},
			},
			"language_code": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book_language).Language_code, nil
				},
			},
			"language_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book_language).Language_name, nil
				},
			},
		},
	})
	countryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Country",
		Fields: graphql.Fields{
			"country_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Country).Country_id, nil
				},
			},
			"country_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Country).Country_name, nil
				},
			},
		},
	})
	cust_orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Cust_order",
		Fields: graphql.Fields{
			"order_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Cust_order).Order_id, nil
				},
			},
			"order_date": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Cust_order).Order_date, nil
				},
			},
			"customer_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Cust_order).Customer_id, nil
				},
			},
			"shipping_method_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Cust_order).Shipping_method_id, nil
				},
			},
			"dest_address_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Cust_order).Dest_address_id, nil
				},
			},
		},
	})
	customerType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"customer_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer).Customer_id, nil
				},
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer).First_name, nil
				},
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer).Last_name, nil
				},
			},
			"email": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer).Email, nil
				},
			},
		},
	})
	customer_addressType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Customer_address",
		Fields: graphql.Fields{
			"customer_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer_address).Customer_id, nil
				},
			},
			"address_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer_address).Address_id, nil
				},
			},
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer_address).Status_id, nil
				},
			},
		},
	})
	order_historyType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order_history",
		Fields: graphql.Fields{
			"history_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_history).History_id, nil
				},
			},
			"order_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_history).Order_id, nil
				},
			},
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_history).Status_id, nil
				},
			},
			"status_date": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_history).Status_date, nil
				},
			},
		},
	})
	order_lineType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order_line",
		Fields: graphql.Fields{
			"line_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_line).Line_id, nil
				},
			},
			"order_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_line).Order_id, nil
				},
			},
			"book_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_line).Book_id, nil
				},
			},
			"price": &graphql.Field{
				Type: scalar.Numeric,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_line).Price, nil
				},
			},
		},
	})
	order_statusType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order_status",
		Fields: graphql.Fields{
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_status).Status_id, nil
				},
			},
			"status_value": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Order_status).Status_value, nil
				},
			},
		},
	})
	publisherType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Publisher",
		Fields: graphql.Fields{
			"publisher_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Publisher).Publisher_id, nil
				},
			},
			"publisher_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Publisher).Publisher_name, nil
				},
			},
		},
	})
	shipping_methodType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Shipping_method",
		Fields: graphql.Fields{
			"method_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Shipping_method).Method_id, nil
				},
			},
			"method_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Shipping_method).Method_name, nil
				},
			},
			"cost": &graphql.Field{
				Type: scalar.Numeric,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Shipping_method).Cost, nil
				},
			},
		},
	})

	customer_addressaddressLoader := batcher.NewLoader[int, *Address](
		pq,
		func(v *Address) int {
			return v.Address_id
		},
		"select address_id,street_number,street_name,city,country_id from address where address_id = any($1)",
	)
	customer_addressType.AddFieldConfig("address", &graphql.Field{
		Type: addressType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customer_addressaddressLoader.Load(p.Context, p.Source.(*Customer_address).Address_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	addresscustomer_addressLoader := batcher.NewListLoader[int, *Customer_address](
		pq,
		func(v *Customer_address) int {
			return v.Address_id
		},
		"select customer_id,address_id,status_id from customer_address where address_id = any($1)",
	)
	addressType.AddFieldConfig("fk_ca_addr", &graphql.Field{
		Type: graphql.NewList(customer_addressType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := addresscustomer_addressLoader.Load(p.Context, p.Source.(*Address).Address_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	cust_orderaddressLoader := batcher.NewLoader[int, *Address](
		pq,
		func(v *Address) int {
			return v.Address_id
		},
		"select address_id,street_number,street_name,city,country_id from address where address_id = any($1)",
	)
	cust_orderType.AddFieldConfig("address", &graphql.Field{
		Type: addressType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := cust_orderaddressLoader.Load(p.Context, p.Source.(*Cust_order).Dest_address_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	addresscust_orderLoader := batcher.NewListLoader[int, *Cust_order](
		pq,
		func(v *Cust_order) int {
			return v.Dest_address_id
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where dest_address_id = any($1)",
	)
	addressType.AddFieldConfig("fk_order_addr", &graphql.Field{
		Type: graphql.NewList(cust_orderType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := addresscust_orderLoader.Load(p.Context, p.Source.(*Address).Address_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	book_authorauthorLoader := batcher.NewLoader[int, *Author](
		pq,
		func(v *Author) int {
			return v.Author_id
		},
		"select author_id,author_name from author where author_id = any($1)",
	)
	book_authorType.AddFieldConfig("author", &graphql.Field{
		Type: authorType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := book_authorauthorLoader.Load(p.Context, p.Source.(*Book_author).Author_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	authorbook_authorLoader := batcher.NewListLoader[int, *Book_author](
		pq,
		func(v *Book_author) int {
			return v.Author_id
		},
		"select book_id,author_id from book_author where author_id = any($1)",
	)
	authorType.AddFieldConfig("fk_ba_author", &graphql.Field{
		Type: graphql.NewList(book_authorType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := authorbook_authorLoader.Load(p.Context, p.Source.(*Author).Author_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	book_authorbookLoader := batcher.NewLoader[int, *Book](
		pq,
		func(v *Book) int {
			return v.Book_id
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where book_id = any($1)",
	)
	book_authorType.AddFieldConfig("book", &graphql.Field{
		Type: bookType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := book_authorbookLoader.Load(p.Context, p.Source.(*Book_author).Book_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookbook_authorLoader := batcher.NewListLoader[int, *Book_author](
		pq,
		func(v *Book_author) int {
			return v.Book_id
		},
		"select book_id,author_id from book_author where book_id = any($1)",
	)
	bookType.AddFieldConfig("fk_ba_book", &graphql.Field{
		Type: graphql.NewList(book_authorType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookbook_authorLoader.Load(p.Context, p.Source.(*Book).Book_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	order_linebookLoader := batcher.NewLoader[int, *Book](
		pq,
		func(v *Book) int {
			return v.Book_id
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where book_id = any($1)",
	)
	order_lineType.AddFieldConfig("book", &graphql.Field{
		Type: bookType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := order_linebookLoader.Load(p.Context, p.Source.(*Order_line).Book_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookorder_lineLoader := batcher.NewListLoader[int, *Order_line](
		pq,
		func(v *Order_line) int {
			return v.Book_id
		},
		"select line_id,order_id,book_id,price from order_line where book_id = any($1)",
	)
	bookType.AddFieldConfig("fk_ol_book", &graphql.Field{
		Type: graphql.NewList(order_lineType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookorder_lineLoader.Load(p.Context, p.Source.(*Book).Book_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookbook_languageLoader := batcher.NewLoader[int, *Book_language](
		pq,
		func(v *Book_language) int {
			return v.Language_id
		},
		"select language_id,language_code,language_name from book_language where language_id = any($1)",
	)
	bookType.AddFieldConfig("book_language", &graphql.Field{
		Type: book_languageType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookbook_languageLoader.Load(p.Context, p.Source.(*Book).Language_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	book_languagebookLoader := batcher.NewListLoader[int, *Book](
		pq,
		func(v *Book) int {
			return v.Language_id
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where language_id = any($1)",
	)
	book_languageType.AddFieldConfig("fk_book_lang", &graphql.Field{
		Type: graphql.NewList(bookType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := book_languagebookLoader.Load(p.Context, p.Source.(*Book_language).Language_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	addresscountryLoader := batcher.NewLoader[int, *Country](
		pq,
		func(v *Country) int {
			return v.Country_id
		},
		"select country_id,country_name from country where country_id = any($1)",
	)
	addressType.AddFieldConfig("country", &graphql.Field{
		Type: countryType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := addresscountryLoader.Load(p.Context, p.Source.(*Address).Country_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	countryaddressLoader := batcher.NewListLoader[int, *Address](
		pq,
		func(v *Address) int {
			return v.Country_id
		},
		"select address_id,street_number,street_name,city,country_id from address where country_id = any($1)",
	)
	countryType.AddFieldConfig("fk_addr_ctry", &graphql.Field{
		Type: graphql.NewList(addressType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := countryaddressLoader.Load(p.Context, p.Source.(*Country).Country_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	order_linecust_orderLoader := batcher.NewLoader[int, *Cust_order](
		pq,
		func(v *Cust_order) int {
			return v.Order_id
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where order_id = any($1)",
	)
	order_lineType.AddFieldConfig("cust_order", &graphql.Field{
		Type: cust_orderType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := order_linecust_orderLoader.Load(p.Context, p.Source.(*Order_line).Order_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	cust_orderorder_lineLoader := batcher.NewListLoader[int, *Order_line](
		pq,
		func(v *Order_line) int {
			return v.Order_id
		},
		"select line_id,order_id,book_id,price from order_line where order_id = any($1)",
	)
	cust_orderType.AddFieldConfig("fk_ol_order", &graphql.Field{
		Type: graphql.NewList(order_lineType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := cust_orderorder_lineLoader.Load(p.Context, p.Source.(*Cust_order).Order_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	order_historycust_orderLoader := batcher.NewLoader[int, *Cust_order](
		pq,
		func(v *Cust_order) int {
			return v.Order_id
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where order_id = any($1)",
	)
	order_historyType.AddFieldConfig("cust_order", &graphql.Field{
		Type: cust_orderType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := order_historycust_orderLoader.Load(p.Context, p.Source.(*Order_history).Order_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	cust_orderorder_historyLoader := batcher.NewListLoader[int, *Order_history](
		pq,
		func(v *Order_history) int {
			return v.Order_id
		},
		"select history_id,order_id,status_id,status_date from order_history where order_id = any($1)",
	)
	cust_orderType.AddFieldConfig("fk_oh_order", &graphql.Field{
		Type: graphql.NewList(order_historyType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := cust_orderorder_historyLoader.Load(p.Context, p.Source.(*Cust_order).Order_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	customer_addresscustomerLoader := batcher.NewLoader[int, *Customer](
		pq,
		func(v *Customer) int {
			return v.Customer_id
		},
		"select customer_id,first_name,last_name,email from customer where customer_id = any($1)",
	)
	customer_addressType.AddFieldConfig("customer", &graphql.Field{
		Type: customerType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customer_addresscustomerLoader.Load(p.Context, p.Source.(*Customer_address).Customer_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	customercustomer_addressLoader := batcher.NewListLoader[int, *Customer_address](
		pq,
		func(v *Customer_address) int {
			return v.Customer_id
		},
		"select customer_id,address_id,status_id from customer_address where customer_id = any($1)",
	)
	customerType.AddFieldConfig("fk_ca_cust", &graphql.Field{
		Type: graphql.NewList(customer_addressType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customercustomer_addressLoader.Load(p.Context, p.Source.(*Customer).Customer_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	cust_ordercustomerLoader := batcher.NewLoader[int, *Customer](
		pq,
		func(v *Customer) int {
			return v.Customer_id
		},
		"select customer_id,first_name,last_name,email from customer where customer_id = any($1)",
	)
	cust_orderType.AddFieldConfig("customer", &graphql.Field{
		Type: customerType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := cust_ordercustomerLoader.Load(p.Context, p.Source.(*Cust_order).Customer_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	customercust_orderLoader := batcher.NewListLoader[int, *Cust_order](
		pq,
		func(v *Cust_order) int {
			return v.Customer_id
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where customer_id = any($1)",
	)
	customerType.AddFieldConfig("fk_order_cust", &graphql.Field{
		Type: graphql.NewList(cust_orderType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customercust_orderLoader.Load(p.Context, p.Source.(*Customer).Customer_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	order_historyorder_statusLoader := batcher.NewLoader[int, *Order_status](
		pq,
		func(v *Order_status) int {
			return v.Status_id
		},
		"select status_id,status_value from order_status where status_id = any($1)",
	)
	order_historyType.AddFieldConfig("order_status", &graphql.Field{
		Type: order_statusType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := order_historyorder_statusLoader.Load(p.Context, p.Source.(*Order_history).Status_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	order_statusorder_historyLoader := batcher.NewListLoader[int, *Order_history](
		pq,
		func(v *Order_history) int {
			return v.Status_id
		},
		"select history_id,order_id,status_id,status_date from order_history where status_id = any($1)",
	)
	order_statusType.AddFieldConfig("fk_oh_status", &graphql.Field{
		Type: graphql.NewList(order_historyType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := order_statusorder_historyLoader.Load(p.Context, p.Source.(*Order_status).Status_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookpublisherLoader := batcher.NewLoader[int, *Publisher](
		pq,
		func(v *Publisher) int {
			return v.Publisher_id
		},
		"select publisher_id,publisher_name from publisher where publisher_id = any($1)",
	)
	bookType.AddFieldConfig("publisher", &graphql.Field{
		Type: publisherType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookpublisherLoader.Load(p.Context, p.Source.(*Book).Publisher_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	publisherbookLoader := batcher.NewListLoader[int, *Book](
		pq,
		func(v *Book) int {
			return v.Publisher_id
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where publisher_id = any($1)",
	)
	publisherType.AddFieldConfig("fk_book_pub", &graphql.Field{
		Type: graphql.NewList(bookType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := publisherbookLoader.Load(p.Context, p.Source.(*Publisher).Publisher_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	cust_ordershipping_methodLoader := batcher.NewLoader[int, *Shipping_method](
		pq,
		func(v *Shipping_method) int {
			return v.Method_id
		},
		"select method_id,method_name,cost from shipping_method where method_id = any($1)",
	)
	cust_orderType.AddFieldConfig("shipping_method", &graphql.Field{
		Type: shipping_methodType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := cust_ordershipping_methodLoader.Load(p.Context, p.Source.(*Cust_order).Shipping_method_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	shipping_methodcust_orderLoader := batcher.NewListLoader[int, *Cust_order](
		pq,
		func(v *Cust_order) int {
			return v.Shipping_method_id
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where shipping_method_id = any($1)",
	)
	shipping_methodType.AddFieldConfig("fk_order_ship", &graphql.Field{
		Type: graphql.NewList(cust_orderType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := shipping_methodcust_orderLoader.Load(p.Context, p.Source.(*Shipping_method).Method_id)
			return func() (any, error) { return thunk() }, nil
		},
	})

	return graphql.Fields{
		"address": &graphql.Field{
			Type:    graphql.NewList(addressType),
			Resolve: batcher.GraphqlField[*Address](pq, "select address_id,street_number,street_name,city,country_id from address"),
		},
		"address_status": &graphql.Field{
			Type:    graphql.NewList(address_statusType),
			Resolve: batcher.GraphqlField[*Address_status](pq, "select status_id,address_status from address_status"),
		},
		"author": &graphql.Field{
			Type:    graphql.NewList(authorType),
			Resolve: batcher.GraphqlField[*Author](pq, "select author_id,author_name from author"),
		},
		"book": &graphql.Field{
			Type:    graphql.NewList(bookType),
			Resolve: batcher.GraphqlField[*Book](pq, "select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book"),
		},
		"book_author": &graphql.Field{
			Type:    graphql.NewList(book_authorType),
			Resolve: batcher.GraphqlField[*Book_author](pq, "select book_id,author_id from book_author"),
		},
		"book_language": &graphql.Field{
			Type:    graphql.NewList(book_languageType),
			Resolve: batcher.GraphqlField[*Book_language](pq, "select language_id,language_code,language_name from book_language"),
		},
		"country": &graphql.Field{
			Type:    graphql.NewList(countryType),
			Resolve: batcher.GraphqlField[*Country](pq, "select country_id,country_name from country"),
		},
		"cust_order": &graphql.Field{
			Type:    graphql.NewList(cust_orderType),
			Resolve: batcher.GraphqlField[*Cust_order](pq, "select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order"),
		},
		"customer": &graphql.Field{
			Type:    graphql.NewList(customerType),
			Resolve: batcher.GraphqlField[*Customer](pq, "select customer_id,first_name,last_name,email from customer"),
		},
		"customer_address": &graphql.Field{
			Type:    graphql.NewList(customer_addressType),
			Resolve: batcher.GraphqlField[*Customer_address](pq, "select customer_id,address_id,status_id from customer_address"),
		},
		"order_history": &graphql.Field{
			Type:    graphql.NewList(order_historyType),
			Resolve: batcher.GraphqlField[*Order_history](pq, "select history_id,order_id,status_id,status_date from order_history"),
		},
		"order_line": &graphql.Field{
			Type:    graphql.NewList(order_lineType),
			Resolve: batcher.GraphqlField[*Order_line](pq, "select line_id,order_id,book_id,price from order_line"),
		},
		"order_status": &graphql.Field{
			Type:    graphql.NewList(order_statusType),
			Resolve: batcher.GraphqlField[*Order_status](pq, "select status_id,status_value from order_status"),
		},
		"publisher": &graphql.Field{
			Type:    graphql.NewList(publisherType),
			Resolve: batcher.GraphqlField[*Publisher](pq, "select publisher_id,publisher_name from publisher"),
		},
		"shipping_method": &graphql.Field{
			Type:    graphql.NewList(shipping_methodType),
			Resolve: batcher.GraphqlField[*Shipping_method](pq, "select method_id,method_name,cost from shipping_method"),
		},
	}
}
