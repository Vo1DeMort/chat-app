import 'package:client_side/login/domain/login_response_model.dart';
import 'package:client_side/login/domain/register_response_model.dart';

sealed class LoginState {}

final class LoginInitialState extends LoginState {}

final class LoginLoading extends LoginState {}

// register related
final class RegisterSuccess extends LoginState {
  final RegisterResponseModel model;

  RegisterSuccess({required this.model});
}

final class RegisterFailed extends LoginState {}

// login related
final class LoginSuccess extends LoginState {
  final LoginResponseModel model;

  LoginSuccess({required this.model});
}

final class LoginFailed extends LoginState {
  final String message;

  LoginFailed({required this.message});
}
