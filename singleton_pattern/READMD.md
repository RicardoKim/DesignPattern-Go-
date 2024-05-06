# Singleton Pattern과 Inversion of Control

## Singleton Pattern

Instance가 생성될때 오직 하나의 Instance가 생성되는 것을 보장하는 것을 목표로 하는 디자인 패턴. 

SingleTon Pattern의 여러 Client들이 접근하는 객체를 하나의 메모리 공간에서 관리할 수 있기 때문에 메모리 효율성이 좋고 객체의 상태관리를 하기에 용이하다.

하지만 이러한 Singleton Pattern은 Application Level에서의 Global State로써 동작하기 때문에 **복잡한 의존성 관리**와 멀티쓰레드 상황에서 **쓰레드 안정성**이 주된 단점으로 뽑히게 된다.

이러한 Singleton Pattern의 단점을 보완하기 위해 적용하는 것이 Inversion of Control 패턴이다.

## Inversion of Control

객체의 의존성 관리 및 주기를 외부 컨테이너등에서 관리하게 하는 디자인 원칙 중에 하나이다.

IoC를 사용하게 되면 위에서 언급한 복잡한 의존성 관리나 Instance 생성시 생기는 Race Condition이 IoC Container에 위임되기 때문에
객체의 생성과 그 객체간의 느슨한 결합으로 이어지게 되고 이에 따라 안정적인 객체의 Life Cycle을 관리할 수 있게 된다.

### IoC에서 객체의 의존성 관리

1. Dependency Injection

객체에 필요한 의존성을 외부에서 주입 받는 방식을 의미한다. 즉 생성자, 메소드, 필드 등을 이용해서 해당 객체의 의존성을 주입하게 된다.

따라서 객체는 자신의 의존성을 직접적으로 관리하지 않고 런타임 환경에서 객체의 설정을 변경할 수 있다.


2. Dependency Lookup

객체가 자신이 필요한 의존성을 의존성 컨테이너나 서비스 레지스트리에 직접적으로 등록하는 방식으로 객체가 스스로 의존성을 관리해야한다.

위 2가지 방법 중에 객체가 자신의 의존성을 직접적으로 관리하게 되면서 생기는 **결합도 증가**, **코드 복잡도 증가**, **테스트의 어려움**으로 인해 Spring Container는 DI방식을 채택하고 있다.

### Spring에서 IoC Container

Spring에서는 객체를 Bean으로 등록하게 되면 Spring Container가 해당 객체의 생성주기를 관리하게 된다.

Spring Container를 이용하는 방식에는 2가지가 있다.

**Bean Factory**

BeanFactory는 Spring IoC 컨테이너 중에 가장 기본적인 형태로 Configuration파일을 읽고 객체를 관리해주는다. 이들은 기본적으로 LazyLoading을 지원하게 된다.


**ApplicationContext**

BeanFactory의 확장된 기능을 제공한다. 빈의 전체 생명 주리를 관리해주는 것 뿐만 아니라 이벤트 전파, 리소스 로딩, AOP지원 등 지원

ApplicationContext
ApplicationContext는 BeanFactory의 확장으로, 보다 다양한 기능을 제공합니다. 

```java
public interface ApplicationContext extends EnvironmentCapable, ListableBeanFactory, HierarchicalBeanFactory,
		MessageSource, ApplicationEventPublisher, ResourcePatternResolver
```

- Bean 관리 : BeanFactory와 같은 방식으로 빈에 대한 설정을 읽고 내부 구조에 저장하게 된다. 이후 Eager Loading을 통해 컨테이너에 빈을 등록함으로써 런타임시 성능을 향상시켜준다.

- 의존성 주입 : 생성자 주입, 새터 주입 등을 통해 DI가 이뤄진다.

- 이벤트 처리 및 AOP 지원 : 이벤트 리스터와 훅을 등록하여 빈의 생명주기 이벤트를 처리하고 AOP지원을 통해 트랜잭션 관리나 보안등의 기능을 제공한다.(ApplicationEventPublisher)

- 리소스 로딩 : 필요한 파일등을 로딩하는 메커니즘을 지원한다.(ResourcePatternResolver)

- 국제화 지원 : 다양한 언어 지원을 위해 메시지 소스를 관리한다.(Message Source)
