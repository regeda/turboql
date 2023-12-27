package bookstore

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stephenafamo/scan/pgxscan"

	"github.com/regeda/turboql/pkg/batcher"
	"github.com/regeda/turboql/pkg/graphqlx/filter"
	"github.com/regeda/turboql/pkg/graphqlx/scalar"
)

type Address struct {
	AddressId    int
	StreetNumber string
	StreetName   string
	City         string
	CountryId    int
}
type AddressStatus struct {
	StatusId      int
	AddressStatus string
}
type Author struct {
	AuthorId   int
	AuthorName string
}
type Book struct {
	BookId          int
	Title           string
	Isbn13          string
	LanguageId      int
	NumPages        int
	PublicationDate time.Time
	PublisherId     int
}
type BookAuthor struct {
	BookId   int
	AuthorId int
}
type BookLanguage struct {
	LanguageId   int
	LanguageCode string
	LanguageName string
}
type Country struct {
	CountryId   int
	CountryName string
}
type CustOrder struct {
	OrderId          int
	OrderDate        time.Time
	CustomerId       int
	ShippingMethodId int
	DestAddressId    int
}
type Customer struct {
	CustomerId int
	FirstName  string
	LastName   string
	Email      string
}
type CustomerAddress struct {
	CustomerId int
	AddressId  int
	StatusId   int
}
type OrderHistory struct {
	HistoryId  int
	OrderId    int
	StatusId   int
	StatusDate time.Time
}
type OrderLine struct {
	LineId  int
	OrderId int
	BookId  int
	Price   pgtype.Numeric
}
type OrderStatus struct {
	StatusId    int
	StatusValue string
}
type Publisher struct {
	PublisherId   int
	PublisherName string
}
type ShippingMethod struct {
	MethodId   int
	MethodName string
	Cost       pgtype.Numeric
}

func NewSchemaConfig(pq pgxscan.Queryer) graphql.SchemaConfig {
	addressType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"address_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).AddressId, nil
				},
			},
			"street_number": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).StreetNumber, nil
				},
			},
			"street_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Address).StreetName, nil
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
					return p.Source.(*Address).CountryId, nil
				},
			},
		},
	})
	addressStatusType := graphql.NewObject(graphql.ObjectConfig{
		Name: "AddressStatus",
		Fields: graphql.Fields{
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*AddressStatus).StatusId, nil
				},
			},
			"address_status": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*AddressStatus).AddressStatus, nil
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
					return p.Source.(*Author).AuthorId, nil
				},
			},
			"author_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Author).AuthorName, nil
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
					return p.Source.(*Book).BookId, nil
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
					return p.Source.(*Book).LanguageId, nil
				},
			},
			"num_pages": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).NumPages, nil
				},
			},
			"publication_date": &graphql.Field{
				Type: scalar.Date,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).PublicationDate, nil
				},
			},
			"publisher_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Book).PublisherId, nil
				},
			},
		},
	})
	bookAuthorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "BookAuthor",
		Fields: graphql.Fields{
			"book_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*BookAuthor).BookId, nil
				},
			},
			"author_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*BookAuthor).AuthorId, nil
				},
			},
		},
	})
	bookLanguageType := graphql.NewObject(graphql.ObjectConfig{
		Name: "BookLanguage",
		Fields: graphql.Fields{
			"language_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*BookLanguage).LanguageId, nil
				},
			},
			"language_code": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*BookLanguage).LanguageCode, nil
				},
			},
			"language_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*BookLanguage).LanguageName, nil
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
					return p.Source.(*Country).CountryId, nil
				},
			},
			"country_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Country).CountryName, nil
				},
			},
		},
	})
	custOrderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "CustOrder",
		Fields: graphql.Fields{
			"order_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustOrder).OrderId, nil
				},
			},
			"order_date": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustOrder).OrderDate, nil
				},
			},
			"customer_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustOrder).CustomerId, nil
				},
			},
			"shipping_method_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustOrder).ShippingMethodId, nil
				},
			},
			"dest_address_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustOrder).DestAddressId, nil
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
					return p.Source.(*Customer).CustomerId, nil
				},
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer).FirstName, nil
				},
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Customer).LastName, nil
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
	customerAddressType := graphql.NewObject(graphql.ObjectConfig{
		Name: "CustomerAddress",
		Fields: graphql.Fields{
			"customer_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustomerAddress).CustomerId, nil
				},
			},
			"address_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustomerAddress).AddressId, nil
				},
			},
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*CustomerAddress).StatusId, nil
				},
			},
		},
	})
	orderHistoryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "OrderHistory",
		Fields: graphql.Fields{
			"history_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderHistory).HistoryId, nil
				},
			},
			"order_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderHistory).OrderId, nil
				},
			},
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderHistory).StatusId, nil
				},
			},
			"status_date": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderHistory).StatusDate, nil
				},
			},
		},
	})
	orderLineType := graphql.NewObject(graphql.ObjectConfig{
		Name: "OrderLine",
		Fields: graphql.Fields{
			"line_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderLine).LineId, nil
				},
			},
			"order_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderLine).OrderId, nil
				},
			},
			"book_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderLine).BookId, nil
				},
			},
			"price": &graphql.Field{
				Type: scalar.Numeric,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderLine).Price, nil
				},
			},
		},
	})
	orderStatusType := graphql.NewObject(graphql.ObjectConfig{
		Name: "OrderStatus",
		Fields: graphql.Fields{
			"status_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderStatus).StatusId, nil
				},
			},
			"status_value": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*OrderStatus).StatusValue, nil
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
					return p.Source.(*Publisher).PublisherId, nil
				},
			},
			"publisher_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*Publisher).PublisherName, nil
				},
			},
		},
	})
	shippingMethodType := graphql.NewObject(graphql.ObjectConfig{
		Name: "ShippingMethod",
		Fields: graphql.Fields{
			"method_id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*ShippingMethod).MethodId, nil
				},
			},
			"method_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*ShippingMethod).MethodName, nil
				},
			},
			"cost": &graphql.Field{
				Type: scalar.Numeric,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return p.Source.(*ShippingMethod).Cost, nil
				},
			},
		},
	})
	addressFilter := filter.NewCursorInput("AddressFilter", graphql.InputObjectConfigFieldMap{
		"address_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"street_number": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"street_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"city": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"country_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	addressStatusFilter := filter.NewCursorInput("AddressStatusFilter", graphql.InputObjectConfigFieldMap{
		"status_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"address_status": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	authorFilter := filter.NewCursorInput("AuthorFilter", graphql.InputObjectConfigFieldMap{
		"author_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"author_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	bookFilter := filter.NewCursorInput("BookFilter", graphql.InputObjectConfigFieldMap{
		"book_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"isbn13": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"language_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"num_pages": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"publisher_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	bookAuthorFilter := filter.NewCursorInput("BookAuthorFilter", graphql.InputObjectConfigFieldMap{
		"book_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"author_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	bookLanguageFilter := filter.NewCursorInput("BookLanguageFilter", graphql.InputObjectConfigFieldMap{
		"language_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"language_code": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"language_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	countryFilter := filter.NewCursorInput("CountryFilter", graphql.InputObjectConfigFieldMap{
		"country_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"country_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	custOrderFilter := filter.NewCursorInput("CustOrderFilter", graphql.InputObjectConfigFieldMap{
		"order_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"customer_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"shipping_method_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"dest_address_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	customerFilter := filter.NewCursorInput("CustomerFilter", graphql.InputObjectConfigFieldMap{
		"customer_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"first_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"last_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	customerAddressFilter := filter.NewCursorInput("CustomerAddressFilter", graphql.InputObjectConfigFieldMap{
		"customer_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"address_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"status_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	orderHistoryFilter := filter.NewCursorInput("OrderHistoryFilter", graphql.InputObjectConfigFieldMap{
		"history_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"order_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"status_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	orderLineFilter := filter.NewCursorInput("OrderLineFilter", graphql.InputObjectConfigFieldMap{
		"line_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"order_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"book_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
	})
	orderStatusFilter := filter.NewCursorInput("OrderStatusFilter", graphql.InputObjectConfigFieldMap{
		"status_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"status_value": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	publisherFilter := filter.NewCursorInput("PublisherFilter", graphql.InputObjectConfigFieldMap{
		"publisher_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"publisher_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})
	shippingMethodFilter := filter.NewCursorInput("ShippingMethodFilter", graphql.InputObjectConfigFieldMap{
		"method_id": &graphql.InputObjectFieldConfig{
			Type: filter.Int,
		},
		"method_name": &graphql.InputObjectFieldConfig{
			Type: filter.String,
		},
	})

	customerAddressAddressLoader := batcher.NewLoader(
		pq,
		func(v *Address) int {
			return v.AddressId
		},
		"select address_id,street_number,street_name,city,country_id from address where 1=1 and address_id = any($1)",
	)
	customerAddressType.AddFieldConfig("address", &graphql.Field{
		Type: addressType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customerAddressAddressLoader.Load(p.Context, p.Source.(*CustomerAddress).AddressId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	addressCustomerAddressLoader := batcher.NewListLoader(
		pq,
		func(v *CustomerAddress) int {
			return v.AddressId
		},
		"select customer_id,address_id,status_id from customer_address where 1=1 and address_id = any($1)",
	)
	addressType.AddFieldConfig("fk_ca_addr", &graphql.Field{
		Type: graphql.NewList(customerAddressType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := addressCustomerAddressLoader.Load(p.Context, p.Source.(*Address).AddressId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	custOrderAddressLoader := batcher.NewLoader(
		pq,
		func(v *Address) int {
			return v.AddressId
		},
		"select address_id,street_number,street_name,city,country_id from address where 1=1 and address_id = any($1)",
	)
	custOrderType.AddFieldConfig("address", &graphql.Field{
		Type: addressType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := custOrderAddressLoader.Load(p.Context, p.Source.(*CustOrder).DestAddressId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	addressCustOrderLoader := batcher.NewListLoader(
		pq,
		func(v *CustOrder) int {
			return v.DestAddressId
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where 1=1 and dest_address_id = any($1)",
	)
	addressType.AddFieldConfig("fk_order_addr", &graphql.Field{
		Type: graphql.NewList(custOrderType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := addressCustOrderLoader.Load(p.Context, p.Source.(*Address).AddressId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookAuthorAuthorLoader := batcher.NewLoader(
		pq,
		func(v *Author) int {
			return v.AuthorId
		},
		"select author_id,author_name from author where 1=1 and author_id = any($1)",
	)
	bookAuthorType.AddFieldConfig("author", &graphql.Field{
		Type: authorType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookAuthorAuthorLoader.Load(p.Context, p.Source.(*BookAuthor).AuthorId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	authorBookAuthorLoader := batcher.NewListLoader(
		pq,
		func(v *BookAuthor) int {
			return v.AuthorId
		},
		"select book_id,author_id from book_author where 1=1 and author_id = any($1)",
	)
	authorType.AddFieldConfig("fk_ba_author", &graphql.Field{
		Type: graphql.NewList(bookAuthorType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := authorBookAuthorLoader.Load(p.Context, p.Source.(*Author).AuthorId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookAuthorBookLoader := batcher.NewLoader(
		pq,
		func(v *Book) int {
			return v.BookId
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where 1=1 and book_id = any($1)",
	)
	bookAuthorType.AddFieldConfig("book", &graphql.Field{
		Type: bookType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookAuthorBookLoader.Load(p.Context, p.Source.(*BookAuthor).BookId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookBookAuthorLoader := batcher.NewListLoader(
		pq,
		func(v *BookAuthor) int {
			return v.BookId
		},
		"select book_id,author_id from book_author where 1=1 and book_id = any($1)",
	)
	bookType.AddFieldConfig("fk_ba_book", &graphql.Field{
		Type: graphql.NewList(bookAuthorType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookBookAuthorLoader.Load(p.Context, p.Source.(*Book).BookId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	orderLineBookLoader := batcher.NewLoader(
		pq,
		func(v *Book) int {
			return v.BookId
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where 1=1 and book_id = any($1)",
	)
	orderLineType.AddFieldConfig("book", &graphql.Field{
		Type: bookType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := orderLineBookLoader.Load(p.Context, p.Source.(*OrderLine).BookId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookOrderLineLoader := batcher.NewListLoader(
		pq,
		func(v *OrderLine) int {
			return v.BookId
		},
		"select line_id,order_id,book_id,price from order_line where 1=1 and book_id = any($1)",
	)
	bookType.AddFieldConfig("fk_ol_book", &graphql.Field{
		Type: graphql.NewList(orderLineType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookOrderLineLoader.Load(p.Context, p.Source.(*Book).BookId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookBookLanguageLoader := batcher.NewLoader(
		pq,
		func(v *BookLanguage) int {
			return v.LanguageId
		},
		"select language_id,language_code,language_name from book_language where 1=1 and language_id = any($1)",
	)
	bookType.AddFieldConfig("book_language", &graphql.Field{
		Type: bookLanguageType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookBookLanguageLoader.Load(p.Context, p.Source.(*Book).LanguageId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookLanguageBookLoader := batcher.NewListLoader(
		pq,
		func(v *Book) int {
			return v.LanguageId
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where 1=1 and language_id = any($1)",
	)
	bookLanguageType.AddFieldConfig("fk_book_lang", &graphql.Field{
		Type: graphql.NewList(bookType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookLanguageBookLoader.Load(p.Context, p.Source.(*BookLanguage).LanguageId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	addressCountryLoader := batcher.NewLoader(
		pq,
		func(v *Country) int {
			return v.CountryId
		},
		"select country_id,country_name from country where 1=1 and country_id = any($1)",
	)
	addressType.AddFieldConfig("country", &graphql.Field{
		Type: countryType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := addressCountryLoader.Load(p.Context, p.Source.(*Address).CountryId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	countryAddressLoader := batcher.NewListLoader(
		pq,
		func(v *Address) int {
			return v.CountryId
		},
		"select address_id,street_number,street_name,city,country_id from address where 1=1 and country_id = any($1)",
	)
	countryType.AddFieldConfig("fk_addr_ctry", &graphql.Field{
		Type: graphql.NewList(addressType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := countryAddressLoader.Load(p.Context, p.Source.(*Country).CountryId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	orderLineCustOrderLoader := batcher.NewLoader(
		pq,
		func(v *CustOrder) int {
			return v.OrderId
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where 1=1 and order_id = any($1)",
	)
	orderLineType.AddFieldConfig("cust_order", &graphql.Field{
		Type: custOrderType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := orderLineCustOrderLoader.Load(p.Context, p.Source.(*OrderLine).OrderId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	custOrderOrderLineLoader := batcher.NewListLoader(
		pq,
		func(v *OrderLine) int {
			return v.OrderId
		},
		"select line_id,order_id,book_id,price from order_line where 1=1 and order_id = any($1)",
	)
	custOrderType.AddFieldConfig("fk_ol_order", &graphql.Field{
		Type: graphql.NewList(orderLineType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := custOrderOrderLineLoader.Load(p.Context, p.Source.(*CustOrder).OrderId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	orderHistoryCustOrderLoader := batcher.NewLoader(
		pq,
		func(v *CustOrder) int {
			return v.OrderId
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where 1=1 and order_id = any($1)",
	)
	orderHistoryType.AddFieldConfig("cust_order", &graphql.Field{
		Type: custOrderType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := orderHistoryCustOrderLoader.Load(p.Context, p.Source.(*OrderHistory).OrderId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	custOrderOrderHistoryLoader := batcher.NewListLoader(
		pq,
		func(v *OrderHistory) int {
			return v.OrderId
		},
		"select history_id,order_id,status_id,status_date from order_history where 1=1 and order_id = any($1)",
	)
	custOrderType.AddFieldConfig("fk_oh_order", &graphql.Field{
		Type: graphql.NewList(orderHistoryType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := custOrderOrderHistoryLoader.Load(p.Context, p.Source.(*CustOrder).OrderId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	customerAddressCustomerLoader := batcher.NewLoader(
		pq,
		func(v *Customer) int {
			return v.CustomerId
		},
		"select customer_id,first_name,last_name,email from customer where 1=1 and customer_id = any($1)",
	)
	customerAddressType.AddFieldConfig("customer", &graphql.Field{
		Type: customerType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customerAddressCustomerLoader.Load(p.Context, p.Source.(*CustomerAddress).CustomerId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	customerCustomerAddressLoader := batcher.NewListLoader(
		pq,
		func(v *CustomerAddress) int {
			return v.CustomerId
		},
		"select customer_id,address_id,status_id from customer_address where 1=1 and customer_id = any($1)",
	)
	customerType.AddFieldConfig("fk_ca_cust", &graphql.Field{
		Type: graphql.NewList(customerAddressType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customerCustomerAddressLoader.Load(p.Context, p.Source.(*Customer).CustomerId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	custOrderCustomerLoader := batcher.NewLoader(
		pq,
		func(v *Customer) int {
			return v.CustomerId
		},
		"select customer_id,first_name,last_name,email from customer where 1=1 and customer_id = any($1)",
	)
	custOrderType.AddFieldConfig("customer", &graphql.Field{
		Type: customerType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := custOrderCustomerLoader.Load(p.Context, p.Source.(*CustOrder).CustomerId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	customerCustOrderLoader := batcher.NewListLoader(
		pq,
		func(v *CustOrder) int {
			return v.CustomerId
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where 1=1 and customer_id = any($1)",
	)
	customerType.AddFieldConfig("fk_order_cust", &graphql.Field{
		Type: graphql.NewList(custOrderType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := customerCustOrderLoader.Load(p.Context, p.Source.(*Customer).CustomerId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	orderHistoryOrderStatusLoader := batcher.NewLoader(
		pq,
		func(v *OrderStatus) int {
			return v.StatusId
		},
		"select status_id,status_value from order_status where 1=1 and status_id = any($1)",
	)
	orderHistoryType.AddFieldConfig("order_status", &graphql.Field{
		Type: orderStatusType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := orderHistoryOrderStatusLoader.Load(p.Context, p.Source.(*OrderHistory).StatusId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	orderStatusOrderHistoryLoader := batcher.NewListLoader(
		pq,
		func(v *OrderHistory) int {
			return v.StatusId
		},
		"select history_id,order_id,status_id,status_date from order_history where 1=1 and status_id = any($1)",
	)
	orderStatusType.AddFieldConfig("fk_oh_status", &graphql.Field{
		Type: graphql.NewList(orderHistoryType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := orderStatusOrderHistoryLoader.Load(p.Context, p.Source.(*OrderStatus).StatusId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	bookPublisherLoader := batcher.NewLoader(
		pq,
		func(v *Publisher) int {
			return v.PublisherId
		},
		"select publisher_id,publisher_name from publisher where 1=1 and publisher_id = any($1)",
	)
	bookType.AddFieldConfig("publisher", &graphql.Field{
		Type: publisherType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := bookPublisherLoader.Load(p.Context, p.Source.(*Book).PublisherId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	publisherBookLoader := batcher.NewListLoader(
		pq,
		func(v *Book) int {
			return v.PublisherId
		},
		"select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where 1=1 and publisher_id = any($1)",
	)
	publisherType.AddFieldConfig("fk_book_pub", &graphql.Field{
		Type: graphql.NewList(bookType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := publisherBookLoader.Load(p.Context, p.Source.(*Publisher).PublisherId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	custOrderShippingMethodLoader := batcher.NewLoader(
		pq,
		func(v *ShippingMethod) int {
			return v.MethodId
		},
		"select method_id,method_name,cost from shipping_method where 1=1 and method_id = any($1)",
	)
	custOrderType.AddFieldConfig("shipping_method", &graphql.Field{
		Type: shippingMethodType,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := custOrderShippingMethodLoader.Load(p.Context, p.Source.(*CustOrder).ShippingMethodId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	shippingMethodCustOrderLoader := batcher.NewListLoader(
		pq,
		func(v *CustOrder) int {
			return v.ShippingMethodId
		},
		"select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where 1=1 and shipping_method_id = any($1)",
	)
	shippingMethodType.AddFieldConfig("fk_order_ship", &graphql.Field{
		Type: graphql.NewList(custOrderType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			thunk := shippingMethodCustOrderLoader.Load(p.Context, p.Source.(*ShippingMethod).MethodId)
			return func() (any, error) { return thunk() }, nil
		},
	})

	return graphql.SchemaConfig{
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{

				"createAddress": &graphql.Field{
					Type: addressType,
					Args: graphql.FieldConfigArgument{
						"address_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"street_number": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"street_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"city": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"country_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: batcher.GraphqlOne[*Address](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into address(address_id,street_number,street_name,city,country_id)values($1,$2,$3,$4,$5)returning address_id,street_number,street_name,city,country_id", []any{p.Args["address_id"], p.Args["street_number"], p.Args["street_name"], p.Args["city"], p.Args["country_id"]}
					}),
				},

				"createAddressStatus": &graphql.Field{
					Type: addressStatusType,
					Args: graphql.FieldConfigArgument{
						"status_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"address_status": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*AddressStatus](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into address_status(status_id,address_status)values($1,$2)returning status_id,address_status", []any{p.Args["status_id"], p.Args["address_status"]}
					}),
				},

				"createAuthor": &graphql.Field{
					Type: authorType,
					Args: graphql.FieldConfigArgument{
						"author_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"author_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*Author](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into author(author_id,author_name)values($1,$2)returning author_id,author_name", []any{p.Args["author_id"], p.Args["author_name"]}
					}),
				},

				"createBook": &graphql.Field{
					Type: bookType,
					Args: graphql.FieldConfigArgument{
						"book_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"title": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"isbn13": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"language_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"num_pages": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"publication_date": &graphql.ArgumentConfig{
							Type: scalar.Date,
						},

						"publisher_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: batcher.GraphqlOne[*Book](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into book(book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id)values($1,$2,$3,$4,$5,$6,$7)returning book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id", []any{p.Args["book_id"], p.Args["title"], p.Args["isbn13"], p.Args["language_id"], p.Args["num_pages"], p.Args["publication_date"], p.Args["publisher_id"]}
					}),
				},

				"createBookAuthor": &graphql.Field{
					Type: bookAuthorType,
					Args: graphql.FieldConfigArgument{
						"book_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"author_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: batcher.GraphqlOne[*BookAuthor](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into book_author(book_id,author_id)values($1,$2)returning book_id,author_id", []any{p.Args["book_id"], p.Args["author_id"]}
					}),
				},

				"createBookLanguage": &graphql.Field{
					Type: bookLanguageType,
					Args: graphql.FieldConfigArgument{
						"language_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"language_code": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"language_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*BookLanguage](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into book_language(language_id,language_code,language_name)values($1,$2,$3)returning language_id,language_code,language_name", []any{p.Args["language_id"], p.Args["language_code"], p.Args["language_name"]}
					}),
				},

				"createCountry": &graphql.Field{
					Type: countryType,
					Args: graphql.FieldConfigArgument{
						"country_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"country_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*Country](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into country(country_id,country_name)values($1,$2)returning country_id,country_name", []any{p.Args["country_id"], p.Args["country_name"]}
					}),
				},

				"createCustOrder": &graphql.Field{
					Type: custOrderType,
					Args: graphql.FieldConfigArgument{
						"order_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"order_date": &graphql.ArgumentConfig{
							Type: graphql.DateTime,
						},

						"customer_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"shipping_method_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"dest_address_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: batcher.GraphqlOne[*CustOrder](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into cust_order(order_id,order_date,customer_id,shipping_method_id,dest_address_id)values($1,$2,$3,$4,$5)returning order_id,order_date,customer_id,shipping_method_id,dest_address_id", []any{p.Args["order_id"], p.Args["order_date"], p.Args["customer_id"], p.Args["shipping_method_id"], p.Args["dest_address_id"]}
					}),
				},

				"createCustomer": &graphql.Field{
					Type: customerType,
					Args: graphql.FieldConfigArgument{
						"customer_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"first_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"last_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*Customer](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into customer(customer_id,first_name,last_name,email)values($1,$2,$3,$4)returning customer_id,first_name,last_name,email", []any{p.Args["customer_id"], p.Args["first_name"], p.Args["last_name"], p.Args["email"]}
					}),
				},

				"createCustomerAddress": &graphql.Field{
					Type: customerAddressType,
					Args: graphql.FieldConfigArgument{
						"customer_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"address_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"status_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: batcher.GraphqlOne[*CustomerAddress](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into customer_address(customer_id,address_id,status_id)values($1,$2,$3)returning customer_id,address_id,status_id", []any{p.Args["customer_id"], p.Args["address_id"], p.Args["status_id"]}
					}),
				},

				"createOrderHistory": &graphql.Field{
					Type: orderHistoryType,
					Args: graphql.FieldConfigArgument{
						"history_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"order_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"status_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"status_date": &graphql.ArgumentConfig{
							Type: graphql.DateTime,
						},
					},
					Resolve: batcher.GraphqlOne[*OrderHistory](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into order_history(history_id,order_id,status_id,status_date)values($1,$2,$3,$4)returning history_id,order_id,status_id,status_date", []any{p.Args["history_id"], p.Args["order_id"], p.Args["status_id"], p.Args["status_date"]}
					}),
				},

				"createOrderLine": &graphql.Field{
					Type: orderLineType,
					Args: graphql.FieldConfigArgument{
						"line_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"order_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"book_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},

						"price": &graphql.ArgumentConfig{
							Type: scalar.Numeric,
						},
					},
					Resolve: batcher.GraphqlOne[*OrderLine](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into order_line(line_id,order_id,book_id,price)values($1,$2,$3,$4)returning line_id,order_id,book_id,price", []any{p.Args["line_id"], p.Args["order_id"], p.Args["book_id"], p.Args["price"]}
					}),
				},

				"createOrderStatus": &graphql.Field{
					Type: orderStatusType,
					Args: graphql.FieldConfigArgument{
						"status_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"status_value": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*OrderStatus](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into order_status(status_id,status_value)values($1,$2)returning status_id,status_value", []any{p.Args["status_id"], p.Args["status_value"]}
					}),
				},

				"createPublisher": &graphql.Field{
					Type: publisherType,
					Args: graphql.FieldConfigArgument{
						"publisher_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"publisher_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: batcher.GraphqlOne[*Publisher](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into publisher(publisher_id,publisher_name)values($1,$2)returning publisher_id,publisher_name", []any{p.Args["publisher_id"], p.Args["publisher_name"]}
					}),
				},

				"createShippingMethod": &graphql.Field{
					Type: shippingMethodType,
					Args: graphql.FieldConfigArgument{
						"method_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},

						"method_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},

						"cost": &graphql.ArgumentConfig{
							Type: scalar.Numeric,
						},
					},
					Resolve: batcher.GraphqlOne[*ShippingMethod](pq, func(p graphql.ResolveParams) (string, []any) {
						return "insert into shipping_method(method_id,method_name,cost)values($1,$2,$3)returning method_id,method_name,cost", []any{p.Args["method_id"], p.Args["method_name"], p.Args["cost"]}
					}),
				},
			},
		}),
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"address": &graphql.Field{
					Type: graphql.NewList(addressType),
					Args: addressFilter,
					Resolve: batcher.GraphqlAll[*Address](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select address_id,street_number,street_name,city,country_id from address where 1=1", nil, p)
					}),
				},
				"address_status": &graphql.Field{
					Type: graphql.NewList(addressStatusType),
					Args: addressStatusFilter,
					Resolve: batcher.GraphqlAll[*AddressStatus](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select status_id,address_status from address_status where 1=1", nil, p)
					}),
				},
				"author": &graphql.Field{
					Type: graphql.NewList(authorType),
					Args: authorFilter,
					Resolve: batcher.GraphqlAll[*Author](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select author_id,author_name from author where 1=1", nil, p)
					}),
				},
				"book": &graphql.Field{
					Type: graphql.NewList(bookType),
					Args: bookFilter,
					Resolve: batcher.GraphqlAll[*Book](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select book_id,title,isbn13,language_id,num_pages,publication_date,publisher_id from book where 1=1", nil, p)
					}),
				},
				"book_author": &graphql.Field{
					Type: graphql.NewList(bookAuthorType),
					Args: bookAuthorFilter,
					Resolve: batcher.GraphqlAll[*BookAuthor](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select book_id,author_id from book_author where 1=1", nil, p)
					}),
				},
				"book_language": &graphql.Field{
					Type: graphql.NewList(bookLanguageType),
					Args: bookLanguageFilter,
					Resolve: batcher.GraphqlAll[*BookLanguage](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select language_id,language_code,language_name from book_language where 1=1", nil, p)
					}),
				},
				"country": &graphql.Field{
					Type: graphql.NewList(countryType),
					Args: countryFilter,
					Resolve: batcher.GraphqlAll[*Country](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select country_id,country_name from country where 1=1", nil, p)
					}),
				},
				"cust_order": &graphql.Field{
					Type: graphql.NewList(custOrderType),
					Args: custOrderFilter,
					Resolve: batcher.GraphqlAll[*CustOrder](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select order_id,order_date,customer_id,shipping_method_id,dest_address_id from cust_order where 1=1", nil, p)
					}),
				},
				"customer": &graphql.Field{
					Type: graphql.NewList(customerType),
					Args: customerFilter,
					Resolve: batcher.GraphqlAll[*Customer](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select customer_id,first_name,last_name,email from customer where 1=1", nil, p)
					}),
				},
				"customer_address": &graphql.Field{
					Type: graphql.NewList(customerAddressType),
					Args: customerAddressFilter,
					Resolve: batcher.GraphqlAll[*CustomerAddress](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select customer_id,address_id,status_id from customer_address where 1=1", nil, p)
					}),
				},
				"order_history": &graphql.Field{
					Type: graphql.NewList(orderHistoryType),
					Args: orderHistoryFilter,
					Resolve: batcher.GraphqlAll[*OrderHistory](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select history_id,order_id,status_id,status_date from order_history where 1=1", nil, p)
					}),
				},
				"order_line": &graphql.Field{
					Type: graphql.NewList(orderLineType),
					Args: orderLineFilter,
					Resolve: batcher.GraphqlAll[*OrderLine](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select line_id,order_id,book_id,price from order_line where 1=1", nil, p)
					}),
				},
				"order_status": &graphql.Field{
					Type: graphql.NewList(orderStatusType),
					Args: orderStatusFilter,
					Resolve: batcher.GraphqlAll[*OrderStatus](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select status_id,status_value from order_status where 1=1", nil, p)
					}),
				},
				"publisher": &graphql.Field{
					Type: graphql.NewList(publisherType),
					Args: publisherFilter,
					Resolve: batcher.GraphqlAll[*Publisher](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select publisher_id,publisher_name from publisher where 1=1", nil, p)
					}),
				},
				"shipping_method": &graphql.Field{
					Type: graphql.NewList(shippingMethodType),
					Args: shippingMethodFilter,
					Resolve: batcher.GraphqlAll[*ShippingMethod](pq, func(p graphql.ResolveParams) (string, []any) {
						return filter.SQL("select method_id,method_name,cost from shipping_method where 1=1", nil, p)
					}),
				},
			},
		}),
	}
}
