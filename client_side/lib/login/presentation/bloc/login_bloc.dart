import 'package:client_side/login/presentation/bloc/login_event.dart';
import 'package:client_side/login/presentation/bloc/login_state.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../data/login_repo.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState> {
  LoginBloc() : super(LoginInitialState()) {
    final repo = LoginRepo();

    on<Register>((event, emit) async {
      final data = repo.register(userName: '', email: '', password: '');
    });

    on<Login>((event, emit) async {
      final data = repo.login(email: '', password: '');
    });
  }
}
