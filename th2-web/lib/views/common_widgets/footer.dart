import 'package:flutter/material.dart';
import 'package:th2_web/configs/color_config.dart';
import 'package:th2_web/configs/text_config.dart';

class Footer extends StatelessWidget {
  const Footer({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.transparent,
      child: Text(
        "© th2system.com, Inc. All rights reserved.",
        style: TextConfigs.kText24_12,
      ),
    );
  }
}
