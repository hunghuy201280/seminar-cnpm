import 'package:hive/hive.dart';
import 'package:injectable/injectable.dart';
import 'package:th2_web/constants/hive_configs.dart';

@module
abstract class HiveModule {
  @Named(HiveConfigs.kConfig)
  Box getConfigPrefBox() => Hive.box(HiveConfigs.kConfig);

  @Named(HiveConfigs.kWallet)
  Box getWalletPrefBox() => Hive.box(HiveConfigs.kWallet);
}
