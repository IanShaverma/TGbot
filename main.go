// Domains

/*
Publication {publicationId channelTelegramID photo text melody}
User {telegramId auth_key_id balance}
Channel {channelTelegramId userTelegramID listening botAdmin}
???Telegram {api}
Prompt {question photo melody}
*/

// UseCases
  //by User
type UserUseCase interface {
	RegistrateUser(userTelegramID domain.User) error
	LinkMyChannel(userTelegramID domain.User, channelTelegramID domain.Channel) error
	TopUpBalance(userTelegramID domain.User balance domain.User) error
    TelegramNewPostUser(photo domain.Publication, userTelegramID domain.User) error
    ListenMyChannel(userTelegramID domain.User, channelTelegramId domain.Channel, listening domain.Channel) error
}
  //by TelegramAPI
type TelegramUseCase interface{
    TelegramNewPostAPI(publicationId domain.Publication) error
}

// Services
type UserService interface {
      //User
    RegistrateUser(userTelegramID domain.User) (domain.User error)
    DeleteUser(userTelegramID domain.User) (domain.User error)
    TopUpUserBalance(userTelegramID domain.User) (domain.User error)
      //Channel
    RegistrateChannel(userTelegramID domain.User, channelTelegramId domain.Channel) (domain.User error, domain.Channel error)
    DeleteChannel(channelTelegramId domain.Channel) (domain.Channel error)
    StartListening(channelTelegramId domain.Channel) (domain.Channel error)
    StopListening(channelTelegramId domain.Channel) (domain.Channel error)
      //Publication
    GeneratePublicationByUser(photo domain.Publication, userTelegramID domain.User) (domain.Publication error, domain.User error)
}

type TelegramService interface {
    //User
  AuthorizeUser(userTelegramID domain.User, auth_key_id domain.User) (domain.User error)
    //Publication
  GeneratePublicationByTelegramAPI(publicationId domain.Publication) (domain.Publication error) //Generate publication
  SetPublicationTelegramAPI(photo domain.Publication, melody domain.Publication, text domain.Publication, channelTelegramID domain.Channel) (error)//Post publication to channel
  SetPublicationUser(photo domain.Publication, melody domain.Publication, text domain.Publication, userTelegramID domain.User) (error) //Post publication privately to user
    //Channel
  GetAdminList(channelTelegramID domain.Channel) (domain.Channel error)
}


// Adapters
type BalanceAPI interface{
    //User
  GetBalance(userTelegramID domain.User) (error)
  SetBalance(balance domain.User) (error)
}
type ListeningChannel interface{
    //Channel
  SetListening(channelTelegramID domain.Channel) (error)
  GetListening(channelTelegramID domain.Channel) (error)// 'true' if channel is listened by Bot
}

type PublicationAPI interface{
    //Publication
  GetPhoto(publicationId domain.Publication, channelTelegramID domain.Channel) (error)
  SetPhoto(photo domain.Publication, publicationId domain.Publication, channelTelegramID domain.Channel) (error)
  GetMelody()
  SetMelody(melody domain.Publication, publicationId domain.Publication, channelTelegramID domain.Channel) (error)
}

type GPTGeneration interface{
    //Prompt
  GenerateChatGPTPrompt(photo domain.Prompt) (error) //Convert photo --> text
  MusicGPTPrompt(question domain.Prompt) (error)//Convert text --> melody
}
