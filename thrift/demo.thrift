namespace go thrift.book

enum BizCode{
   SUCCESS=200,
   INTERNAL_ERROR=500
}

exception BizException{
 1: string msg,
 2: BizCode code
}


struct BookInfo{
 1: required i32 bookId
 2: string bookName
}

struct Book{
 1: required i32 bookId,
 2: BookInfo info
}


service BookService{
   BookInfo getBookInfo(1: i32 bookId);
   list<Book> getAllBook();
}
