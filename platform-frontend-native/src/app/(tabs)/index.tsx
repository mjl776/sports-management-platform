import { View, Text, ScrollView, StyleSheet } from 'react-native';
export default function DashboardScreen() {
  return (
    <ScrollView style={styles.container}>
    <View style={styles.header}>
      <Text style={styles.title}>Sports Management</Text>
      <Text style={styles.subtitle}>Dashboard</Text>
    </View>
  </ScrollView>
  );
}

const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#F9FAFB',
    },
    header: {
      padding: 24,
      backgroundColor: '#1E40AF',
    },
    title: {
      fontSize: 28,
      fontWeight: 'bold',
      color: 'white',
    },
    subtitle: {
      fontSize: 16,
      color: '#BFDBFE',
      marginTop: 4,
    },
  });