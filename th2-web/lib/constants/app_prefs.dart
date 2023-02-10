import 'package:injectable/injectable.dart';
import 'package:th2_web/models/prefs/configspref.dart';

import '../models/prefs/walletpref.dart';

@singleton
class AppPref {
  final ConfigsPref config;
  final WalletPref wallet;

  const AppPref({
    required this.config,
    required this.wallet,
  });
}
