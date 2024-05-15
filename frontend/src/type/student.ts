export interface Student {
    id : string
    first_name : string
    last_name : string
    surname : string
    passport_seria : number
    passport_number : number
    birth_date : string
    email : string
    password : string
    country : string
    city : string
    street : number
    house_number : number
    apartment_number : number
    enroll_year : string
    specialization : string
    course : number
    group : number
    link : string
}

export interface StudentShort {
    id : string
    first_name : string
    last_name : string
    surname : string
    enroll_year : string
    specialization : string
    course : number
    group : number
}