

//   ((import drivers for use SQLite from git))


package main

    import (
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
    )

 func main() {
        db, err := sql.Open("mysql")
        checkErr(err)


//     ((Insert a new asset and price and count))


stmt, err := db.Prepare("INSERT assets SET id=?,Assets Name=?,Unit Price=?,Count=?")
        checkErr(err)

        res, err := stmt.Exec("3","Apartment", "800 M", "1")
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println(id)


//  ((Update three Price rows))


stmt, err = db.Prepare("UPDATE assets SET Price=? where Id=?")
        checkErr(err)

        res, err = stmt.Exec("1,300 M", 1)
        checkErr(err)

        res, err = stmt.Exec("40 K", 2)
        checkErr(err)	

        res, err = stmt.Exec("200 M", 3)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

        fmt.Println(affect)


//  ((Delete Two assets (Dollar and Apartment) and Insert cash = '900 M'))


stmt, err = db.Prepare("DELETE from assets where id=?")
        checkErr(err)

        res, err = stmt.Exec(2)
        checkErr(err)
	  
        res, err = stmt.Exec(3)
        checkErr(err)	

        affect, err = res.RowsAffected()
        checkErr(err)
        
stmt, err := db.Prepare("INSERT assets SET id=?,Assets Name=?,Unit Price=?,Count=?")
        checkErr(err)

        res, err := stmt.Exec("2","Cash", "40 K", "22500")
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        affect, err = res.RowsAffected()
        checkErr(err)

        fmt.Println(affect)


//  ((Insert two assets(ETH & EOS = '620 M') and Update Cash = 280 M '))



stmt, err = db.Prepare("UPDATE assets SET Price=? where Id=?")
        checkErr(err)

res, err = stmt.Exec("280 M", 2)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

stmt, err := db.Prepare("INSERT assets SET id=?,Assets Name=?,Unit Price=?,Count=?")
        checkErr(err)

        res, err := stmt.Exec("3","EOS", "40 K", "5 K")
        checkErr(err)

        res, err := stmt.Exec("4","ETH", "60 M", "7")
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        db.Close()


//  ((Set a function for error respond))


}

    func checkErr(err error) 
       {
        if err != nil {panic(err)}
        }
}