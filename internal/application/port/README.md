# 포트(Port) 레이어

- Hexagonal Architecture에서 착안한 아이디어로, 의존성 역전(DIP)을 실현하는 중요한 역할을 담당합니다.

- 파일들에서 확인할 수 있듯이 본 프로젝트에서 필요한 모든 로직들을 interface로 정의한 것을 확인할 수 있습니다. 이를 통해 마치 전자제품을 아무거나 사도 포트가 같기 때문에 충전을 할 수 있듯이, port 레이어 덕분에 확장성과 유지보수성애 매우 증가합니다.

![예시 이미지](/docs/images/port_and_adapter_example.png)

만약 더 자세한 정보를 원한다면 [hexagonal-startup](https://github.com/fullgukbap/hexagonal-startup) 리포지토리를 참고해주세요.
