package bookeeping

func NewFirebase(path string,project string) (*Firebase,error) {

    return newFirebase(path,project)
}

func NewCache() (*Cache) {
    
    return newCache()
}
