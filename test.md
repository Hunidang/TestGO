# FMS - Main.go 

## 1. Data

```Go
int test string
fmt.Println(test)
```

```uml

start

=>start: 시작
end=>end
o1=>operation: 오퍼레이션1
o2=>operation: 오퍼레이션2
o3=>operation: 오퍼레이션3
c1=>condition: 2 or 3 ?

start->o1->c1
c1(yes)->o2->end
c1(no)->o3->end

```

## 2. Test SQL Code 
```SQL
select * from user
where 1=1
and key = 1234;
```