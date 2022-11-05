package THARUNK;
import java.util.Arrays;
public class Bublesort {
            public static void main(String[] args) {
                int[] arr ={5,3,45,2,3 };
                buble(arr);
                System.out.println( Arrays.toString(arr));
            }
            static void buble(int[] arr)
            {
                for (int i = 0; i < arr.length; i++) {
                    boolean swap=false;
                    for (int j = 1; j<arr.length-i; j++) {
                        if(arr[j] < arr[j-1])
                        {
                            int temp=arr[j];
                            arr[j]=arr[j-1];
                            arr[j-1]=temp;
                            swap=true;
                        }
                    }
                    if(swap==false){
                        break;
                    }
                }
            }
        }

