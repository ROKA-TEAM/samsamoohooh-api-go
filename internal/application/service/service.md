# 서비스(Service) 레이어

- Serivce 폴더에서는 도메인의 비지니스 로직을 작성합니다. 이를 통해 비지니스 로직에 집중할 수 있기에, 변경사항이 발생할 경우 해당 로직을 중점적으로 개선할 수 있어 유지보수성에 매우 뛰어납니다.

- 또한 각각의 서비스객체는 Port의 Repository interface를 주입받아서 사용하기 떄문에, 유연함과 확장성에 매우 뛰어납니다.

- [port 코드 확인하러가기](https://github.com/ROKA-TEAM/samsamoohooh-go-api/tree/pinned/2024_10_28/internal/application/port)
