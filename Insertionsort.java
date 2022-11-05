package THARUNK;

import java.util.Arrays;

public class Insertionsort {
    public static void main(String[] args) {
      int[] arr={5,3,4,2,1};

        for (int i = 0; i < arr.length; i++) {
            int last = arr.length - i - 1;
            int m= maxis(arr, 0, last);
          swap(arr,m,last);

        }
        System.out.println(Arrays.toString(arr));


    }
   static void swap(int[] arr,int first ,int second) {
            int temp=arr[first];
            arr[first]=arr[second];
            arr[second]=temp;
   }
    static int   maxis(int[] arr,int start,int end)
    {
        int m=start;
        for(int i =start;i<=end;i++)
        {
            if(arr[m]<arr[i]){
                m=i;
            }

        }
        return m;
    }
}

