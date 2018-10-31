# 用Go实现常用设计模式

## 1. 策略模式 (strategy)

**意图:**
定义一系列的算法,把它们一个个封装起来, 并且使它们可相互替换。

**关键代码:**
实现同一个接口

**应用实例:**
1. 主题的更换，每个主题都是一种策略
2. 旅行的出游方式，选择骑自行车、坐汽车，每一种旅行方式都是一个策略。 
3. JAVA AWT 中的 LayoutManager。


## 2. 装饰器模式 (decorator)

**意图:**
装饰器模式动态地将责任附加到对象上。若要扩展功能，装饰者提供了比继承更有弹性的替代方案。

**关键代码:**
装饰器和被装饰对象实现同一个接口，装饰器中使用了被装饰对象

**应用实例:**
1. JAVA中的IO流 
```java
new DataInputStream(new FileInputStream("test.txt");
```
2. 人穿衣服，人则为被装饰对象，衣服为装饰器

## 3. 适配器模式 (adaptor)
> 适配器适合用于解决新旧系统（或新旧接口）之间的兼容问题，而不建议在一开始就直接使用

**意图:**
适配器模式将一个类的接口，转换成客户期望的另一个接口。适配器让原本接口不兼容的类可以合作无间

**关键代码:**
适配器中持有旧接口对象，并实现新接口

**应用实例:**
1. 充电器转接口头
2. 安卓的ListView
```java
ListView lv=(ListView) findViewById(R.id.lv);
String []data ={"hi","nihao","yes","no"};
ArrayAdapter<String> adapter=new ArrayAdapter<>(this,android.R.layout.simple_list_item_1,data);
lv.setAdapter(adapter);
```

## 4. 单例 (singleton)
**意图:**
使程序运行期间只存在一个实例对象

## 5. 观察者模式 (observer)
**意图:**
定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。

**关键代码:**
被观察者持有了集合存放观察者(收通知的为观察者)

**应用实例:**
1. 报纸订阅，报社为被观察者，订阅的人为观察者
2. MVC模式，当model改变时，View视图会自动改变，model为被观察者，View为观察者
