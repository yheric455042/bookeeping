package bookeeping

func GetFirebase(path string,project string) (*Firebase,error) {

    return newFirebase(path,project)
}
