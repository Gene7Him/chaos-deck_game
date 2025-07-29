// React Native App w/ navigation and socket connect
import React, { useEffect, useState } from 'react';
import { View, Text, Button } from 'react-native';

export default function App() {
  const [ws, setWs] = useState(null);
  const [hand, setHand] = useState([]);
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/ws');
    socket.onopen = () => console.log('Connected');
    socket.onmessage = (e) => {
      const data = JSON.parse(e.data);
      if (data.type === 'hand') setHand(data.cards);
      if (data.type === 'chat') setMessages(prev => [...prev, data.message]);
    };
    socket.onerror = console.error;
    socket.onclose = () => console.log('Disconnected');
    setWs(socket);
  }, []);

  const playCard = (card) => {
    if (ws) ws.send(JSON.stringify({ type: 'play_card', data: card }));
  };

  return (
    <View style={{ padding: 20 }}>
      <Text>CHAOS DECK</Text>
      {hand.map((c, i) => (
        <Button key={i} title={`${c.color} ${c.value}`} onPress={() => playCard(c)} />
      ))}
    </View>
  );
}
