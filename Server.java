import java.net.*;
import java.nio.charset.StandardCharsets;

class Server
{
    public static void main(String args[]) throws Exception
    {
        DatagramSocket serverSocket = new DatagramSocket(8081);
        byte[] receiveData = new byte[1024];
        while(true)
        {
            DatagramPacket receivePacket = new DatagramPacket(receiveData, receiveData.length);
            serverSocket.receive(receivePacket);
            String s = "Message Received: " + new String( receivePacket.getData()).trim();
            System.out.println(s);
            serverSocket.send(new DatagramPacket(s.getBytes(), s.length(), receivePacket.getSocketAddress()));
        }
    }
}