package clock

import "fmt"

type Clock struct{
    hours, minutes int
}
// Define the Clock type here.

func New(h, m int) Clock {
    if m<0{
        h=h%24-((m*-1)/60)-1
        m=60-(m*-1)%60
    }else{
    	h=h%24+m/60
    	m=m%60
    }
    if m==60{
        m-=60
        h+=1
    }
    h=h%24
    if h<0{
        h=24-(h*-1)
    }
    var c Clock=Clock{h,m}
	return c
}

func (c Clock) Add(m int) Clock {
	minutes:=c.minutes+m
    c.hours+=minutes/60
    c.minutes=minutes%60
    c.hours=c.hours%24
	return c
}

func (c Clock) Subtract(m int) Clock {
	t:=m/60
    c.hours-=t+1
    c.minutes-=m%60
    c.minutes+=60
    if c.minutes >=60{
        c.minutes-=60
        c.hours+=1
    }
    if c.hours<0{
        c.hours=24-(c.hours*-1)
    }
    c.hours=c.hours%24
    return c
}

func (c Clock) String() string {
    for ; c.hours<0;{
        c.hours+=24
    }
	return fmt.Sprintf("%02d:%02d",c.hours%24,c.minutes)
}
