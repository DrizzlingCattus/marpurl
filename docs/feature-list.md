## 임시 설계 기록

- 업로드 할 수 있는 프론트엔드가 필요함
  - 로그인 화면~ 처음 화면
    - 세션기능 필요
  - 업로드 화면 ~ 로그인 후 실질적인 작업이 이루어지는 곳
    - 올릴 대상(폴더로 고정하자)을 정할 수 있다.
    - 올릴 대상의 이름을 정할 수 있다. (실질적)
      - 이름이 이전에 업로드한 파일과 겹칠시에 다시 입력하라는 중복체크 명령
	- [ ]이를 위한 백엔드 api
    - 결과 url을 보여준다. (결과 https://[server-host]/marpurl/[ppt-name])
    - 이미 올라간 ppt 파일들의 목록을 볼 수 있다. (pagination)
      - 이를 위한 백엔드 api

- 백엔드
  - 프론트엔드 파일을 배포할 API
  - 세션, 로그인 API
  - DB 연결
    - [x] gorm과 mysql 연결
    - [x] 연결 로직을 따로 분리 (/db)
  - 마이그레이션
    - [x] 모델 변화를 즉시 마이그레이션
    - [] 업데이트마다 마이그레이션 백업 파일 생성
  - Model 설계 및 생성
    - [x] PPT 모델 생성
    - [x] Model api들을 따로 분리 (/model)
  - 받을 파일을 저장할 공간
    - 이름 겹치지않게 저장
  - marp-cli에 빌드 명령 (shell 명령 사용 방법 및 완료 시점 체크)
  - 환경 변수 분리
    - [x] 개발, 프로덕션, 테스트 환경마다 환경 변수 분리할 수 있게 구성
    - [] 인자로 어떤 환경인지 세팅할 수 있다.
    - [x] env pkg로 분리
    - [] 프로젝트 root path를 얻는 방법
    - [] 테스트 코드 작성
  - 테스트 환경 구성
    - [ ] dbcleaner pkg
    - [x] testfixture
      - [] fixture load가 각 테스트마다 호출시에 어떻게 동작하는지 알아볼것
    - [] testify
      - [x] suite
      - [x] assert
      - [] mock


업로드 Rest API

- [ ] DB 
