//TODO: implement theme for app

// by default is to follow the system
// this is api from sdk
// thememode.light , dark system

import 'package:flutter/material.dart';

ThemeData buildAppTheme(ThemeMode mode) {
  switch (mode) {
    case ThemeMode.light:
      return ThemeData.light();

    case ThemeMode.dark:
      return ThemeData.dark();

    case ThemeMode.system:
      return ThemeData.light();
  }
}
