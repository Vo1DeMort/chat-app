sealed class LoginEvent {}

final class Login extends LoginEvent {
  final String email;
  final String password;

  Login({required this.email, required this.password});
}

final class Register extends LoginEvent {
  final String userName;
  final String email;
  final String password;

  Register({
    required this.userName,
    required this.email,
    required this.password,
  });
}
