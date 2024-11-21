// Domains

/*
Publication {publicationId channelId photo text melody}
User {telegramId auth_key_id balance}
Channel {channelTelegramId userTelegramID listening botAdmin}
???Telegram {api}
Prompt {question photo}
*/

// UseCases
type UserUseCase interface {
	RegistrateUser(userTelegramID domain.User) error
	LinkMyChannel(userTelegramID domain.User, channelid domain.Channel) error
	Accounting(userTelegramID domain.User balance domain.User) error
    TelegramNewPostUser(photo domain.Publication, userTelegramID domain.User) error
    ListenMyChannel(userTelegramID domain.User, channelTelegramId domain.Channel, listening domain.Channel) error
    TelegramNewPostAPI(publicationId domain.Publication) error
}

// Services
type UserService interface {
    //User
    RegistrateUser()
    AuthorizeUser()
    DeleteUser()
    TopUpBalance()
    //Channel
    RegistrateChannel()
    DeleteChannel()
    StartListening()
    StopListening()
    //Publication
    GetPhoto()
    SetPhoto()
    GetMelody()
    SetMelody()
 
}

// Adapters
type BalanceAPI interface{
  //User
  GetBalance()
  SetBalance()
}
type 
  //Channel
  GetAdminList()
  //Prompt
  ChatGPTPrompt()
  MusicGPTPrompt()
}