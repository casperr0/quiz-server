from django.test import TestCase

from qa_server.models import Player


class EnumParseTestCase(TestCase):
    def test_player_parse_platform(self):

        self.assertEqual(Player.parse_platform("messenGER"), Player.Platform.MESSENGER)
        self.assertEqual(Player.parse_platform("TelegRam"), Player.Platform.TELEGRAM)
        self.assertEqual(Player.parse_platform("DISCORD"), Player.Platform.DISCORD)
        self.assertEqual(Player.parse_platform("NETCAT"), Player.Platform.NETCAT)
        self.assertEqual(Player.parse_platform("Line"), Player.Platform.LINE)
        self.assertEqual(Player.parse_platform("mewe"), Player.Platform.MEWE)
